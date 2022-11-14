package grpc

import (
	"api_gateway/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

func init_grpc_client() (*grpc.ClientConn, proto.WeatherForecastServiceClient, context.Context, context.CancelFunc) {
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

func GetForecast(forecastInput proto.ForecastInput) *proto.ForecastOutput {
	conn, c, ctx, cancel := init_grpc_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetForecast(ctx, &forecastInput)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r
}
