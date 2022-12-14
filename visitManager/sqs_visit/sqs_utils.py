import time
from pprint import pprint

import boto3 as boto3


def sendInviteRequestMessage(queue_url, creator, username, visit, partecipation):



    sqs = boto3.client('sqs', region_name='us-east-1')

    # Send message to SQS queue
    sqs.send_message(
        QueueUrl=queue_url,
        DelaySeconds=10,
        MessageAttributes={
            'InviteRequest_'+username: {
                'DataType': 'String',
                'StringValue': str(partecipation)
            },
            'Creator': {
                'DataType': 'String',
                'StringValue': creator
            },
            'Visit': {
                'DataType': 'String',
                'StringValue': visit
            },
        },
        MessageBody=(
            "Request message"
        )

    )

def removeRequestMessage(queue_url, username, visit):

    sqs = boto3.client('sqs', region_name='us-east-1')

    # Send message to SQS queue
    while True:
        response = sqs.receive_message(
            QueueUrl=queue_url,
            AttributeNames=[
                'All'
            ],
            MaxNumberOfMessages=1,
            MessageAttributeNames=[
                'InviteRequest_'+username,
                'Visit'
            ],
            VisibilityTimeout=3,
            WaitTimeSeconds=0
        )
        if "Messages" not in response:
            break

        msg = response['Messages'][0]
        if "MessageAttributes" in msg:
            msgattrs = msg.get("MessageAttributes")
            participants = msgattrs.get("InviteRequest_"+username)
            if not participants:
                continue
            print(visit)
            visitM = msgattrs.get("Visit").get("StringValue")
            print(visitM)
            if visit==visitM:
                receipt_handle = msg['ReceiptHandle']

                # Delete received message from queue
                sqs.delete_message(
                    QueueUrl=queue_url,
                    ReceiptHandle=receipt_handle
                )
                print('Received and deleted message: %s' % msg)
                return 1
    return 0
