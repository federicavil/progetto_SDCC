import logging
import time
from concurrent import futures

import boto3
import grpc
import visitManager_pb2
import visitManager_pb2_grpc
import psycopg2
import mysql.connector
from configparser import ConfigParser

from sqs_visit.sqs_utils import sendInviteRequestMessage, removeRequestMessage


"""
    Esegue il parsing del file ini di configurazione accedendo ad una determinata sezione
    Ritorna un dizionario con i campi letti dal file
    @:param parser: config parser del file
    @:param section: sezione alla quale accedere
"""
def parse_config(parser, section):
    conf = {}
    if parser.has_section(section):
        params = parser.items(section)
        for param in params:
            conf[param[0]] = param[1]
    else:
        raise Exception('Section {0} not found in the file'.format(section))
    return conf

"""   
    Ritorna le configurazioni lette dal file in input
    @:param filename: path del file di configurazione
"""
def config(filename='./conf/database.ini'):
    # create a parser
    parser = ConfigParser()
    # read config file
    parser.read(filename)
    section = parse_config(parser, "app_mode")["app_mode"]
    db = parse_config(parser, section)

    # get section, default to postgresql
    return db


class ManageVisitServicer(visitManager_pb2_grpc.ManageVisitServicer):
    """Provides methods that implement functionality of route guide server."""

    """
        Esegue la connessione al database
    """
    def dbConnection(self):
        # read connection parameters
        params = config()
        dbtype = params["dbtype"]
        cond = dbtype == "postgres"
        params.pop("dbtype")
        quote = params["quote"]
        params.pop("quote")
        params.pop("host_port")
        conn = None
        try:
            if cond:
                conn = psycopg2.connect(**params)
            else:
                conn = mysql.connector.connect(**params)
        except (Exception, psycopg2.DatabaseError, mysql.connector.DatabaseError) as error:
            raise Exception
        return conn, quote

    """
        Per prima cosa chiama la funzione che si connette al database, dopodiché inizializza la coda SQS
    """
    def __init__(self):
        """ Connect to the PostgreSQL database server """
        self.conn = None
        i = 1
        while i < 10:
            try:
                self.conn, self.quote = self.dbConnection()
                i = 10
            except:
                i = i+1
                time.sleep(5)
        if self.conn is None:
            return

        sqs = boto3.resource('sqs', region_name='us-east-1')
        #Crea coda se non esiste e ne estrae il link per l'utilizzo
        try:
            queue = sqs.get_queue_by_name(
                QueueName='SDCC_Antonangeli_Villani_Notification',
            )
        except Exception:
            queue = sqs.create_queue(QueueName='SDCC_Antonangeli_Villani_Notification', Attributes={'DelaySeconds': '5'})
        self.queue_url = queue.url


    """
        Implementa il servizio di aggiunta di una nuova visita
        @:param request: richiesta contenente la nuova visita da aggiungere
        @:param context: contesto
    """
    def AddNewVisit(self, request, context):
        cur = self.conn.cursor()
        username = request.Username
        pathname = request.Pathname
        timestamp = request.Timestamp
        from datetime import datetime
        date_format = "%Y-%m-%dT%H:%M"
        ts1 = datetime.strptime(timestamp, date_format)
        import pandas as pd
        ts = pd.Timestamp(ts1)

        quote = self.quote
        sql = """INSERT INTO """ + quote + """Visit""" + quote + """ (""" + quote + """ID_Path"""+quote + """, """ + quote +"""Timestamp"""+quote+""", """+quote+"""Creator"""+quote+""") VALUES ('""" + pathname + """', """ + str(
            ts.to_julian_date()) + """, '""" + username + """')"""
        cur.execute(sql)
        sql = """SELECT max("""+quote+"""ID"""+quote+""") FROM """ + quote + """Visit""" + quote  # WHERE Timestamp"="""+str(ts.to_julian_date())
        cur.execute(sql)
        id_visit = str(cur.fetchone()[0])
        sql = """INSERT INTO """ + quote + """User_to_Visit""" + quote + """ ("""+quote+"""ID_User"""+quote+""", """+quote+"""ID_Visit"""+quote+""", """+quote+"""Accepted"""+quote+""") VALUES ('""" + username + """', """ + id_visit + """, True)"""
        cur.execute(sql)
        self.conn.commit()
        ret = visitManager_pb2.Return(Ret=True)
        return ret
    """
        Implementa il servizio che recupera le informazioni di una visita dato il suo id
        @:param request: richiesta contenente l'id della visita
        @:param context: contesto
    """
    def GetVisitByID(self, request, context):
        cur = self.conn.cursor()
        quote = self.quote
        sql = """ SELECT * FROM """ + quote + """Visit""" + quote +""" WHERE """+quote+"""ID"""+quote+""" = """ + request.Value
        cur.execute(sql)
        try:
            d = cur.fetchall()[0]
        except IndexError:
            return visitManager_pb2.Visit(ID_Visit=0, ID_Path="", Timestamp=-1, Creator = "")
        from datetime import datetime
        date_format = "%Y-%m-%dT%H:%M"
        import pandas as pd
        t = pd.to_datetime(d[2], origin='julian', unit='D')
        ts = t.strftime("%d/%m/%Y, %H:%M")
        sql = """ SELECT * FROM """ + quote + """User_to_Visit"""+quote +""" WHERE """+quote+"""ID_Visit"""+quote+"""=""" + str(d[0]) + """ AND """+quote+"""Accepted"""+quote+""" = true"""
        cur.execute(sql)
        resPart = cur.fetchall()
        participants = []
        for p in resPart:
            participants.append(p[1])
        visit = visitManager_pb2.Visit(ID_Visit=d[0], ID_Path=d[1], Timestamp=ts, Creator = d[3], Participants = participants)
        return visit

    """
        Implementa il servizio che recupera tutte le visite relative ad uno specifico utente
        @:param request: richiesta contenente l'id dell'utente di cui recuperare le visite
        @:param context: contesto
    """
    def GetAllVisits(self, request, context):
        cur = self.conn.cursor()
        quote = self.quote
        sql = """SELECT * FROM """ + quote + """Visit""" + quote + """ JOIN """ + quote + """User_to_Visit""" + quote + """ on """ + quote + """Visit""" + quote + """."""+quote+"""ID"""+quote+"""=""" + quote + """User_to_Visit""" + quote + """.""" +quote + """ID_Visit"""+quote +""" WHERE """+quote+"""ID_User"""+quote+""" LIKE '""" + request.ID.replace("'", "''") + """' AND """+quote+"""Accepted"""+quote+""" = true"""
        cur.execute(sql)
        data = cur.fetchall()

        visits = []
        for d in data:
            sql = """ SELECT * FROM """ + quote + """User_to_Visit"""+quote +""" WHERE """+quote+"""ID_Visit"""+quote+"""=""" + str(d[0]) +""" AND """+quote+"""Accepted"""+quote+""" = true"""
            cur.execute(sql)
            resPart = cur.fetchall()
            participants = []
            for p in resPart:
                participants.append(p[1])
            import pandas as pd
            t = pd.to_datetime(d[2], origin='julian', unit='D')
            ts = t.strftime("%d/%m/%Y, %H:%M")
            visit = visitManager_pb2.Visit(ID_Visit=d[0], ID_Path=d[1], Timestamp=ts, Creator=d[3], Participants = participants)
            visits.append(visit)
        response = visitManager_pb2.Visits(Visit=visits)
        return response

    """
        Implementa il servizio di invito di un utente ad una visita
        @:param request: richiesta contenente l'id della visita e dell'utente da invitare
        @:param context: contesto
    """
    def InviteUserToVisit(self, request, context):

        cur = self.conn.cursor()
        username = request.Username.replace("'", "''")
        id_visit = request.ID_Visit
        quote = self.quote
        sql = """SELECT """+quote+"""Visit"""+quote+"""."""+quote+"""Creator"""+quote+""" FROM """+quote+"""Visit"""+quote+""" WHERE """+quote+"""Visit"""+quote+"""."""+quote+"""ID"""+quote+"""='""" + id_visit + """'"""
        cur.execute(sql)
        creator = str(cur.fetchone()[0])
        ret = visitManager_pb2.Return(Ret=-10)
        if creator == username:
            ret = visitManager_pb2.Return(Ret=-3)
            return ret
        sql = """INSERT INTO """+quote+"""User_to_Visit"""+quote+""" ("""+quote+"""ID_User"""+quote+""", """+quote+"""ID_Visit"""+quote+""") VALUES ('""" + username + """', """ + id_visit + """)"""
        try:
            cur.execute(sql)
        except psycopg2.errors.ForeignKeyViolation:
            cur.execute("ROLLBACK")
            self.conn.commit()
            ret = visitManager_pb2.Return(Ret=-1)
            return ret
        except psycopg2.errors.UniqueViolation:
            cur.execute("ROLLBACK")
            self.conn.commit()
            ret = visitManager_pb2.Return(Ret=-2)
            return ret
        except Exception:
            cur.execute("ROLLBACK")
            self.conn.commit()
            ret = visitManager_pb2.Return(Ret=-2)
            return ret
        sql = """SELECT """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_User"""+quote+""" FROM """+quote+"""User_to_Visit"""+quote+""" WHERE """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""Accepted"""+quote+"""=True AND """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_Visit"""+quote+"""='""" + id_visit + """'"""
        cur.execute(sql)
        participants = cur.fetchall()
        participantList = []
        for p in participants:
            participantList.append(p[0])
        try:
            sendInviteRequestMessage(self.queue_url, creator, username, id_visit, participantList)
            self.conn.commit()
            ret = visitManager_pb2.Return(Ret=1)
        except Exception:
            self.conn.rollback()
            ret = visitManager_pb2.Return(Ret=-4)
        return ret

    """
        Implementa il servizio di accettazione o no di un invito
        @:param request: richiesta contenente l'id della visita, l'id dell'utente e la risposta
        @:param context: contesto
    """
    def AcceptOrRefuseInvite(self, request, context):
        cur = self.conn.cursor()
        username = request.Username.replace("'", "''")
        id_visit = request.ID_Visit
        response = request.Response
        quote = self.quote
        sql = """UPDATE """+quote+"""User_to_Visit"""+quote+""" SET """+quote+"""Accepted"""+quote+"""=""" + response + """ WHERE """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_Visit"""+quote+"""='""" + id_visit + """' AND """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_User"""+quote+"""='""" + username + """'"""
        cur.execute(sql)
        try:
            deleted = removeRequestMessage(self.queue_url, username, id_visit)
            if deleted == 1:
                self.conn.commit()
                ret = visitManager_pb2.Return(Ret=1)
            else:
                self.conn.rollback()
                ret = visitManager_pb2.Return(Ret=0)
        except Exception:
            self.conn.rollback()
            ret = visitManager_pb2.Return(Ret=-1)
        return ret

"""
    Definisce il server grpc tramite il quale è possibile invocare il servizio
"""
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    visitManager_pb2_grpc.add_ManageVisitServicer_to_server(
        ManageVisitServicer(), server)
    configurations = config()
    server.add_insecure_port('[::]:'+configurations["host_port"])
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
