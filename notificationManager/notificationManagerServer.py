import json
import logging
import time
from concurrent import futures

import boto3
import botocore
import grpc
from botocore.exceptions import ClientError

import notificationManager_pb2
import notificationManager_pb2_grpc
import psycopg2
from configparser import ConfigParser

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
def config(filename='./conf/conf.ini'):
    # create a parser
    parser = ConfigParser()
    # read config file
    parser.read(filename)
    section = parse_config(parser, "app_mode")["app_mode"]
    db = parse_config(parser, section)

    # get section, default to postgresql
    return db


class NotificationManagerServicer(notificationManager_pb2_grpc.NotificationManagerServicer):
    """Provides methods that implement functionality of route guide server."""

    """
        Inizializza la coda di messaggi SQS
    """
    def __init__(self):
        self.messages = {}
        self.sqs = boto3.client('sqs', region_name='us-east-1')
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
        Recupera gli inviti relativi ad un utente dalla coda SQS e li restituisce
        @Param: userId: id dell'utente del quale restituire gli inviti
    """
    def receiveInviteRequestMessage(self, userId):
        # Create SQS client
        messages = []
        retValues = []
        while True:
            response = self.sqs.receive_message(
                QueueUrl=self.queue_url,
                AttributeNames=[
                    'All'
                ],
                MaxNumberOfMessages=1,
                MessageAttributeNames=[
                    'InviteRequest_'+userId,
                    'Creator',
                    'Visit'
                ],
                VisibilityTimeout=2,
                WaitTimeSeconds=0
            )
            if "Messages" not in response:
                break
            messages.append(response)
        for response in messages:

            if "MessageAttributes" in response['Messages'][0]:
                msgattrs = response['Messages'][0].get("MessageAttributes")
                participants = msgattrs.get("InviteRequest_"+userId)
                if not participants:
                    continue
                visit = msgattrs.get("Visit")
                creator = msgattrs.get("Creator")
                myJson = {"visit": visit.get("StringValue"), "creator": creator.get("StringValue"), "participants": participants.get("StringValue")}
                retValues.append(myJson)
        return retValues

    """
        Implementa il servizio di restituzione degli inviti associati ad un utente definito nel file proto
        @:param request: richiesta grpc
        @:param context: contesto di chiamata
    """
    def GetInvites(self, request, context):
        userId = request.Username
        ret = notificationManager_pb2.InviteOutput(Invites=json.dumps(self.receiveInviteRequestMessage(userId)))
        return ret

"""
    Definisce il server grpc tramite il quale è possibile invocare il servizio
"""
def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  notificationManager_pb2_grpc.add_NotificationManagerServicer_to_server(
      NotificationManagerServicer(), server)
  configurations = config()
  server.add_insecure_port('[::]:'+configurations["host_port"])
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
  logging.basicConfig()
  serve()