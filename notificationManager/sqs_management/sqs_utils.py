import time
from pprint import pprint

import boto3 as boto3

class NotificationManager():

    def __init__(self):
        self.messages = {}
        self.sqs = boto3.client('sqs', region_name='us-east-1')
        self.queue_url = "https://sqs.us-east-1.amazonaws.com/983687073675/Notification"

    def receiveInviteRequestMessage(self, username):
        # Create SQS client
        messages = []
        while True:
            response = self.sqs.receive_message(
                QueueUrl=self.queue_url,
                AttributeNames=[
                    'InviteRequest_' + username
                ],
                MaxNumberOfMessages=1,
                MessageAttributeNames=[
                    'InviteRequest_' + username
                ],
                VisibilityTimeout=5,
                WaitTimeSeconds=0
            )
            if "Messages" not in response:
                break
            messages.append(response)
            # receipt_handle = response['Messages'][0]['ReceiptHandle']
            #
            # # Delete received message from queue
            # self.sqs.delete_message(
            #     QueueUrl=self.queue_url,
            #     ReceiptHandle=receipt_handle
            # )

        for response in messages:
            if "MessageAttributes" in response['Messages'][0]:
                msgaattrs = response['Messages'][0].get("MessageAttributes")
                msgVisit = msgaattrs.get("InviteRequest_" + username).get("StringValue")
                print(msgVisit)
                # TODO: risposta GRPC a VisitManager
        # notif.freeQueue()

    #
    # def sendInviteResponseMessage(self, username="", visit="", partecipation=""):
    #
    #     # Send message to SQS queue
    #     response = self.sqs.send_message(
    #         QueueUrl=self.queue_url,
    #         DelaySeconds=10,
    #         MessageAttributes={
    #             'InviteResponse_'+username+visit: {
    #                 'DataType': 'String',
    #                 'StringValue': partecipation
    #             },
    #         },
    #         MessageBody=(
    #             ""
    #         )
    #     )
    #
    #     print(response['MessageId'])


if __name__ == "__main__":
    notif = NotificationManager()
    notif.receiveInviteRequestMessage("pippo")
    # notif.sendInviteRequestMessage()
    # notif.receiveInviteRequestMessage("pippi")
