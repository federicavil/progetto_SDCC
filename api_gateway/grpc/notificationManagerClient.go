package grpc

import (
	"api_gateway/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

func init_grpc_notification_manager_client() (*grpc.ClientConn, proto.NotificationManagerClient, context.Context, context.CancelFunc) {
	// Set up a connection to the server.

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewNotificationManagerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return conn, c, ctx, cancel
}

func GetInvites(inviteInput proto.InviteInput) []byte {
	conn, c, ctx, cancel := init_grpc_notification_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetInvites(ctx, &inviteInput)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	invites, _ := json.Marshal(r.Invites)
	return invites
}
