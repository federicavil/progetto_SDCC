package grpc

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultNameVM = "world"
)

func init_grpc_visit_manager_client() (*grpc.ClientConn, proto.ManageVisitClient, context.Context, context.CancelFunc) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(conf.GetConnectionConf("visit_manager"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewManageVisitClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return conn, c, ctx, cancel
}

func AddNewVisit(visit proto.InputVisit) int {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()

	r, err := c.AddNewVisit(ctx, &visit)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(*r.Ret)
}

func GetAllVisits(username string) []byte {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetAllVisits(ctx, &proto.User{
		ID: &username,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	visits, _ := json.Marshal(r.Visit)
	return visits
}

func GetVisitByID(id string) []byte {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetVisitByID(ctx, &proto.ID_Visit{
		Value: &id,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	visit, _ := json.Marshal(r)
	fmt.Println(string(visit))
	return visit
}

func InviteUserToVisit(invite proto.Invite) int {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.InviteUserToVisit(ctx, &invite)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(*r.Ret)
}

func AcceptOrRefuseInvite(invite proto.InviteResponse) int {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.AcceptOrRefuseInvite(ctx, &invite)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(*r.Ret)
}
