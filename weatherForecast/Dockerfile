FROM python:3

COPY . .

RUN pip install grpcio
RUN pip install protobuf
RUN pip install psycopg2-binary
RUN pip install requests

CMD [ "python", "./weatherForecastServer.py" ]