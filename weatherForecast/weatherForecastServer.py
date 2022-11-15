import json
import logging
from concurrent import futures
from pprint import pprint

import grpc
import weatherForecastService_pb2
import weatherForecastService_pb2_grpc
import psycopg2
import requests
from configparser import ConfigParser

class WeatherForecastServiceServicer(weatherForecastService_pb2_grpc.WeatherForecastServiceServicer):
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

    def initt(self):
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

    def GetForecast(self, request, context):
        path = request.Path
        parsedPath = json.loads(path)
        city = parsedPath["City"]
        province = parsedPath["Province"]
        region = parsedPath["Region"]
        dict_keys = ["time","temperature","humidity","precipitation","cloud_cover","wind_speed"]
        time = []
        temperature = []
        humidity = []
        precipitation = []
        cloud_cover = []
        wind_speed = []
        response = requests.get("https://geocoding-api.open-meteo.com/v1/search?name="+city)
        resp = response.content.decode()
        jsn = json.loads(resp)
        try:
            lat=jsn["results"][0]["latitude"]
            long=jsn["results"][0]["longitude"]
            response = requests.get("https://api.open-meteo.com/v1/forecast?latitude="+str(lat)+"&longitude="+str(long)+"&hourly=temperature_2m,relativehumidity_2m,precipitation,cloudcover,windspeed_10m")
            resp = response.content.decode()
            jsn = json.loads(resp)
            #Forecasts: "{\"time\" : [\"11/11/2022\",\"12/11/2022\",\"13/11/2022\"],\"temperature\" : [4,7,12],\"cloud_cover\" : [40,100,75],\"precipitation\" : [9,10,0],\"wind_speed\" : [12,3,5],\"humidity\" : [4,8,2]}",
            for i in range(0,len(jsn["hourly"]["cloudcover"]),24):
                from datetime import datetime
                date_format = "%Y-%m-%dT%H:%M"
                ts1 = datetime.strptime(jsn["hourly"]["time"][i], date_format)
                date_format = "%d/%m/%Y"
                ts2 = datetime.strftime(ts1, date_format)
                time.append(ts2)
                temperature.append(jsn["hourly"]["temperature_2m"][i])
                humidity.append(jsn["hourly"]["relativehumidity_2m"][i])
                precipitation.append(jsn["hourly"]["precipitation"][i])
                cloud_cover.append(jsn["hourly"]["cloudcover"][i])
                wind_speed.append(jsn["hourly"]["windspeed_10m"][i])
        except KeyError:
            pass

        dict_values = [time,temperature,humidity,precipitation,cloud_cover,wind_speed]
        result_dict = {dict_keys[i]: dict_values[i] for i in range(len(dict_values))}
        print(result_dict)
        #GEOCODING: https://geocoding-api.open-meteo.com/v1/search?name=Berlin
        #METEO: https://api.open-meteo.com/v1/forecast?latitude=51.5002&longitude=-0.1262&hourly=temperature_2m,relativehumidity_2m,precipitation,cloudcover,windspeed_10m

        ret = weatherForecastService_pb2.ForecastOutput(Forecasts=str(json.dumps(result_dict)))
        return ret

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    weatherForecastService_pb2_grpc.add_WeatherForecastServiceServicer_to_server(
      WeatherForecastServiceServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()