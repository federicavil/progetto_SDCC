FROM python:3

COPY . .

RUN pip install grpcio
RUN pip install protobuf
RUN pip install psycopg2-binary
RUN pip install mysql-connector-python
RUN pip install pandas
RUN pip install boto3

EXPOSE 9093

CMD [ "python", "./visitManagerServer.py" ]