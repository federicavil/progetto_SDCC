import json
import logging
import time
from concurrent import futures
from pprint import pprint

import boto3
import botocore
import grpc
from botocore.exceptions import ClientError

import notificationManager_pb2
import notificationManager_pb2_grpc
import psycopg2
from configparser import ConfigParser

class NotificationManagerServicer(notificationManager_pb2_grpc.NotificationManagerServicer):
    """Provides methods that implement functionality of route guide server."""

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

    def GetInvites(self, request, context):
        userId = request.Username
        ret = notificationManager_pb2.InviteOutput(Invites=json.dumps(self.receiveInviteRequestMessage(userId)))
        return ret

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  notificationManager_pb2_grpc.add_NotificationManagerServicer_to_server(
      NotificationManagerServicer(), server)
  server.add_insecure_port('[::]:9094')
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
  logging.basicConfig()
  serve()