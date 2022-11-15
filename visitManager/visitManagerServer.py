import logging
from concurrent import futures

import grpc
import visitManager_pb2
import visitManager_pb2_grpc
import psycopg2
from configparser import ConfigParser

class ManageVisitServicer(visitManager_pb2_grpc.ManageVisitServicer):
    """Provides methods that implement functionality of route guide server."""

    # def __init__(self):
    #     self.db = visitManager_resources.visitManager_guide_database()


    def config(self, filename='./conf/database.ini', section='postgresql'):
        # create a parser
        parser = ConfigParser()
        # read config file
        parser.read(filename)

        # get section, default to postgresql
        db = {}
        if parser.has_section(section):
            params = parser.items(section)
            for param in params:
                db[param[0]] = param[1]
        else:
            raise Exception('Section {0} not found in the {1} file'.format(section, filename))

        return db

    def __init__(self):
        """ Connect to the PostgreSQL database server """
        conn = None
        try:
            # read connection parameters
            params = self.config()

            # connect to the PostgreSQL server
            print('Connecting to the PostgreSQL database...')
            self.conn = psycopg2.connect(**params)
            print("Connected")
            # create a cursor
        except (Exception, psycopg2.DatabaseError) as error:
            print(error)

    def AddNewVisit(self, request, context):
        cur = self.conn.cursor()
        username = request.Username
        pathname = request.Pathname
        timestamp = request.Timestamp
        print("AddNewVisit("+username+","+pathname+","+timestamp+")")
        from datetime import datetime
        date_format = "%Y-%m-%dT%H:%M"
        ts1 = datetime.strptime(timestamp, date_format)
        import pandas as pd
        ts = pd.Timestamp(ts1)
        print(ts)
        sql = """INSERT INTO public."Visit" ("ID_Path", "Timestamp", "Creator") VALUES ('"""+pathname+"""', """+str(ts.to_julian_date())+""", '"""+username+"""')"""
        cur.execute(sql)
        sql = """SELECT max("ID") FROM public."Visit" """ #WHERE Timestamp"="""+str(ts.to_julian_date())
        cur.execute(sql)
        id_visit = str(cur.fetchone()[0])
        sql = """INSERT INTO public."User_to_Visit" ("ID_User", "ID_Visit", "Accepted") VALUES ('""" + username + """', """ + id_visit + """, True)"""
        cur.execute(sql)
        self.conn.commit()
        ret = visitManager_pb2.Return(Ret=True)
        return ret

    def GetAllVisits(self, request, context):
        cur = self.conn.cursor()
        sql=""" SELECT * FROM public."Visit" JOIN public."User_to_Visit" on public."Visit"."ID"=public."User_to_Visit"."ID_Visit" WHERE "ID_User" LIKE '"""+request.ID+"'"
        cur.execute(sql)
        data = cur.fetchall()
        visits = []
        for d in data:
            visit = visitManager_pb2.Visit(ID_Visit = d[0], ID_Path=d[1], Timestamp=d[2])
            visits.append(visit)
        response = visitManager_pb2.Visits(Visit=visits)
        return response

    def InviteUserToVisit(self, request, context):

        cur = self.conn.cursor()
        username = request.Username
        id_visit = request.ID_Visit
        print("InviteUserToVisit("+username+","+id_visit+")")
        sql = """SELECT "Visit"."Creator" FROM public."Visit" WHERE "Visit"."ID"='"""+id_visit+"""'"""
        cur.execute(sql)
        creator = str(cur.fetchone()[0])
        if creator == username:
            ret = visitManager_pb2.Return(Ret=-3)
            return ret
        sql = """INSERT INTO public."User_to_Visit" ("ID_User", "ID_Visit") VALUES ('""" + username + """', """ + id_visit + """)"""
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
        ###TODO: CHIAMATA MICROSERVIZIO NOTIFY PER INVIO NOTIFICA ALL'UTENTE
        ret = visitManager_pb2.Return(Ret=1)
        return ret

    def AcceptOrRefuseInvite(self, request, context):
        cur = self.conn.cursor()
        username = request.Username
        id_visit = request.ID_Visit
        response = request.Response
        sql = """UPDATE public."User_to_Visit" SET "User_to_Visit"."Accepted"="""+response+""" WHERE "User_to_Visit"."ID_Visit"='"""+id_visit+"""' AND "User_to_Visit"."Username"='"""+username+"""'"""
        cur.execute(sql)
        self.conn.commit()
        ret = visitManager_pb2.Return(Ret=1)
        return ret

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  visitManager_pb2_grpc.add_ManageVisitServicer_to_server(
      ManageVisitServicer(), server)
  server.add_insecure_port('[::]:50051')
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
  logging.basicConfig()
  serve()