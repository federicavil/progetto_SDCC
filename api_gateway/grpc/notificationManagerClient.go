package grpc

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
* Inizializza le strutture dati usate nella chiamata gRPC al microservizio NotificationManager
* @returns {*grpc.ClientConn}: connessione gRPC al microservizio NotificationManager
* @returns {proto.NotificationManagerClient}: istanza di client generata da file ".proto" con protobuf
* @returns {context.Context}: permette di contattare il server gRPC
* @returns {context.CancelFunc}: usato con defer per terminare il lavoro a fine funzione
* @returns {error}
 */
func init_grpc_notification_manager_client() (*grpc.ClientConn, proto.NotificationManagerClient, context.Context, context.CancelFunc, error) {
	// Set up a connection to the server.
	var c proto.NotificationManagerClient
	var ctx context.Context
	var cancel context.CancelFunc

	conn, err := grpc.Dial(conf.GetConnectionConf("notification"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		//log.Fatalf("did not connect: %v", err)
	} else {

		c = proto.NewNotificationManagerClient(conn)

		// Contact the server and print out its response.
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	}
	return conn, c, ctx, cancel, err
}

/*
* Effettua chiamata gRPC per ricevere lista degli inviti
* @param {proto.InviteInput}: contiene userId dell'utente di cui vogliamo ottenere tutti gli inviti
* @returns {[]byte}: lista di inviti in json
* @returns {error}
 */
func GetInvites(inviteInput proto.InviteInput) ([]byte, error) {
	conn, c, ctx, cancel, err := init_grpc_notification_manager_client()
	var invites []byte
	if err == nil {
		defer conn.Close()
		defer cancel()
		r, er := c.GetInvites(ctx, &inviteInput)
		if er != nil {
			invites, _ = json.Marshal("[]")
			err = er
			fmt.Println("ERROR")
		} else {
			invites, _ = json.Marshal(r.Invites)
		}
	} else {
		invites, _ = json.Marshal("[]")
		fmt.Println("ERROR")
	}
	fmt.Println(string(invites))
	return invites, err
}
