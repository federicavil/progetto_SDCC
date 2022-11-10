package grpc

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func init_grpc_client() (*grpc.ClientConn, ManageVisitClient, context.Context, context.CancelFunc) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := NewManageVisitClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, c, ctx, cancel
}

func AddNewVisit(visit InputVisit) int {
	conn, c, ctx, cancel := init_grpc_client()
	defer conn.Close()
	defer cancel()

	r, err := c.AddNewVisit(ctx, &visit)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(*r.Ret)
}

func GetAllVisits(username string) []byte {
	conn, c, ctx, cancel := init_grpc_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetAllVisits(ctx, &User{
		ID: &username,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	visits, _ := json.Marshal(r.Visit)
	return visits
}

func InviteUserToVisit(invite Invite) int {
	conn, c, ctx, cancel := init_grpc_client()
	defer conn.Close()
	defer cancel()
	r, err := c.InviteUserToVisit(ctx, &invite)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(*r.Ret)
}
