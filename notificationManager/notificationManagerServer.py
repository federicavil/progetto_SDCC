import json
import logging
import time
from concurrent import futures
from pprint import pprint

import boto3
import grpc
import notificationManager_pb2
import notificationManager_pb2_grpc
import psycopg2
from configparser import ConfigParser

class NotificationManagerServicer(notificationManager_pb2_grpc.NotificationManagerServicer):
    """Provides methods that implement functionality of route guide server."""

    def __init__(self):
        self.messages = {}
        self.sqs = boto3.client('sqs', region_name='us-east-1')

        self.queue_url = "https://sqs.us-east-1.amazonaws.com/983687073675/Notification"
        print("SQS queue initialized.")

    def receiveInviteRequestMessage(self, username):
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
                    'InviteRequest_'+username,
                    'Creator',
                    'Visit'
                ],
                VisibilityTimeout=5,
                WaitTimeSeconds=0
            )
            if "Messages" not in response:
                break
            messages.append(response)
        for response in messages:

            if "MessageAttributes" in response['Messages'][0]:
                msgattrs = response['Messages'][0].get("MessageAttributes")
                participants = msgattrs.get("InviteRequest_"+username)
                if not participants:
                    continue
                visit = msgattrs.get("Visit")
                creator = msgattrs.get("Creator")
                myJson = {"visit": visit.get("StringValue"), "creator": creator.get("StringValue"), "participants": participants.get("StringValue")}
                retValues.append(myJson)
        return retValues

    def GetInvites(self, request, context):
        print("Get invites")
        username = request.Username
        ret = notificationManager_pb2.InviteOutput(Invites=json.dumps(self.receiveInviteRequestMessage(username)))
        print(ret)
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