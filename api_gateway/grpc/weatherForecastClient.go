package grpc

import (
	"api_gateway/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
* Inizializza le strutture dati usate nella chiamata gRPC al microservizio WeatherForecast
* @returns {*grpc.ClientConn}: connessione gRPC al microservizio WeatherForecast
* @returns {proto.WeatherForecastServiceClient}: istanza di client generata da file ".proto" con protobuf
* @returns {context.Context}: permette di contattare il server gRPC
* @returns {context.CancelFunc}: usato con defer per terminare il lavoro a fine funzione
* @returns {error}
 */
func init_grpc_weather_forecast_client() (*grpc.ClientConn, proto.WeatherForecastServiceClient, context.Context, context.CancelFunc) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewWeatherForecastServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, c, ctx, cancel
}

/*
* Effettua chiamata gRPC per ottenere informazioni sul meteo
* @param {proto.ForecastInput}: contiene il luogo in cui si trova il percorso
* @returns {*proto.ForecastOutput}: contiene un json con le informazioni del meteo in una settimana
 */
func GetForecast(forecastInput proto.ForecastInput) *proto.ForecastOutput {
	conn, c, ctx, cancel := init_grpc_weather_forecast_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetForecast(ctx, &forecastInput)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r
}
