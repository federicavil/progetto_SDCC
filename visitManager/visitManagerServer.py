import logging
import time
from concurrent import futures

import grpc
import visitManager_pb2
import visitManager_pb2_grpc
import psycopg2
import mysql.connector
from configparser import ConfigParser

from sqs_visit.sqs_utils import sendInviteRequestMessage, removeRequestMessage


class ManageVisitServicer(visitManager_pb2_grpc.ManageVisitServicer):
    """Provides methods that implement functionality of route guide server."""

    # def __init__(self):
    #     self.db = visitManager_resources.visitManager_guide_database()

    def sendInviteRequestMessage(self, username="", visit="", partecipation="", participants="{}"):

        # Send message to SQS queue
        response = self.sqs.send_message(
            QueueUrl=self.queue_url,
            DelaySeconds=10,
            MessageAttributes={
                'InviteResponse_'+username+visit: {
                    'DataType': 'String',
                    'StringValue': partecipation
                },
                'Participants': {
                    'DataType': 'String',
                    'StringValue': participants
                },
            },
            MessageBody=(
                ""
            )
        )

        print(response['MessageId'])

    def parse_config(self, parser, section):
        conf = {}
        if parser.has_section(section):
            params = parser.items(section)
            for param in params:
                conf[param[0]] = param[1]
        else:
            raise Exception('Section {0} not found in the file'.format(section))
        return conf

    def config(self, filename='./conf/database.ini'):
        # create a parser
        parser = ConfigParser()
        # read config file
        parser.read(filename)
        section = self.parse_config(parser, "app_mode")["app_mode"]
        print("SECTION" + section)
        db = self.parse_config(parser, section)

        # get section, default to postgresql
        """ db = {}
        if parser.has_section(section):
            params = parser.items(section)
            for param in params:
                db[param[0]] = param[1]
        else:
            raise Exception('Section {0} not found in the {1} file'.format(section, filename)) """
        print(db)
        return db

    def dbConnection(self):
        # read connection parameters
        params = self.config()
        dbtype = params["dbtype"]
        cond = dbtype == "postgres"
        params.pop("dbtype")
        quote = dbtype = params["quote"]
        params.pop("quote")
        print(params)
        conn = None
        try:
            if cond:
                print("postgres")
                conn = psycopg2.connect(**params)
                print("Connected")
            else:
                print("mysql")
                conn = mysql.connector.connect(**params)
                print("Connected")
        except (Exception, psycopg2.DatabaseError, mysql.connector.DatabaseError) as error:
            print("CONNECTION ERROR")
            print(error)
            raise Exception
        return conn, quote


    def __init__(self):
        """ Connect to the PostgreSQL database server """
        i = 1
        while i < 10:
            try:
                self.conn, self.quote = self.dbConnection()
                i = 10
            except:
                i = i+1
                time.sleep(1)


    def AddNewVisit(self, request, context):
        cur = self.conn.cursor()
        username = request.Username
        pathname = request.Pathname
        timestamp = request.Timestamp
        print("AddNewVisit(" + username + "," + pathname + "," + timestamp + ")")
        from datetime import datetime
        date_format = "%Y-%m-%dT%H:%M"
        ts1 = datetime.strptime(timestamp, date_format)
        import pandas as pd
        ts = pd.Timestamp(ts1)
        print(ts)
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

    def GetAllVisits(self, request, context):
        cur = self.conn.cursor()
        quote = self.quote
        sql = """ SELECT * FROM """ + quote + """Visit""" + quote + """ JOIN """ + quote + """User_to_Visit""" + quote + """ on """ + quote + """Visit""" + quote + """."ID"=""" + quote + """User_to_Visit""" + quote + """.""" +quote + """ID_Visit"""+quote +""" WHERE """+quote+"""ID_User"""+quote+""" LIKE '""" + request.ID + "'"
        cur.execute(sql)
        data = cur.fetchall()
        visits = []
        for d in data:
            visit = visitManager_pb2.Visit(ID_Visit=d[0], ID_Path=d[1], Timestamp=d[2])
            visits.append(visit)
        response = visitManager_pb2.Visits(Visit=visits)
        return response

    def InviteUserToVisit(self, request, context):

        cur = self.conn.cursor()
        username = request.Username
        id_visit = request.ID_Visit
        print("InviteUserToVisit(" + username + "," + id_visit + ")")
        quote = self.quote
        sql = """SELECT """+quote+"""Visit"""+quote+"""."""+quote+"""Creator"""+quote+""" FROM """+quote+"""Visit"""+quote+""" WHERE """+quote+"""Visit"""+quote+"""."""+quote+"""ID"""+quote+"""='""" + id_visit + """'"""
        cur.execute(sql)
        creator = str(cur.fetchone()[0])
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
        self.conn.commit()
        sql = """SELECT """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_User"""+quote+""" FROM """+quote+"""User_to_Visit"""+quote+""" WHERE """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""Accepted"""+quote+"""=True AND """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_Visit"""+quote+"""='""" + id_visit + """'"""
        cur.execute(sql)
        participants = cur.fetchall()
        participantList = []
        for p in participants:
            participantList.append(p[0])
        sendInviteRequestMessage(creator, username, id_visit, participantList)
        ret = visitManager_pb2.Return(Ret=1)
        return ret

    def AcceptOrRefuseInvite(self, request, context):
        cur = self.conn.cursor()
        username = request.Username
        id_visit = request.ID_Visit
        response = request.Response
        quote = self.quote
        sql = """UPDATE """+quote+"""User_to_Visit"""+quote+""" SET """+quote+"""Accepted"""+quote+"""=""" + response + """ WHERE """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_Visit"""+quote+"""='""" + id_visit + """' AND """+quote+"""User_to_Visit"""+quote+"""."""+quote+"""ID_User"""+quote+"""='""" + username + """'"""
        cur.execute(sql)
        deleted = removeRequestMessage(username, id_visit)
        if deleted == 1:
            self.conn.commit()
            ret = visitManager_pb2.Return(Ret=1)
        else:
            self.conn.rollback()
            ret = visitManager_pb2.Return(Ret=0)
        return ret


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    visitManager_pb2_grpc.add_ManageVisitServicer_to_server(
        ManageVisitServicer(), server)
    server.add_insecure_port('[::]:9093')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
