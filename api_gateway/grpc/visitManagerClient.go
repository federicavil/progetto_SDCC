package grpc

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultNameVM = "world"
)

/*
* Inizializza le strutture dati usate nella chiamata gRPC al microservizio VisitManager
* @returns {*grpc.ClientConn}: connessione gRPC al microservizio VisitManager
* @returns {proto.ManageVisitClient}: istanza di client generata da file ".proto" con protobuf
* @returns {context.Context}: permette di contattare il server gRPC
* @returns {context.CancelFunc}: usato con defer per terminare il lavoro a fine funzione
* @returns {error}
 */
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

/*
* Effettua chiamata gRPC per aggiungere una nuova visita al sistema
* @param {proto.InputVisit}: contiene informazioni relative alla visita da aggiungere
* @returns {int}
 */
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

/*
* Effettua chiamata gRPC per ricevere la lista delle visite relative a un userId
* @param {string}: userId dell'utente
* @returns {[]byte}: lista delle visite in json
 */
func GetAllVisits(userId string) []byte {
	conn, c, ctx, cancel := init_grpc_visit_manager_client()
	defer conn.Close()
	defer cancel()
	r, err := c.GetAllVisits(ctx, &proto.User{
		ID: &userId,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	visits, _ := json.Marshal(r.Visit)
	return visits
}

/*
* Effettua chiamata gRPC per ricevere le informazioni relative a una visita dato il suo id
* @param {string}: id della visita
* @returns {[]byte}: informazioni della visita in json
 */
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
	return visit
}

/*
* Effettua chiamata gRPC per invitare un utente a una visita
* @param {proto.Invite}: contiene userId e id della visita
* @returns {int}
 */
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

/*
* Effettua chiamata gRPC per accettare o rifiutare un invito
* @param {proto.InviteResponse}: contiene userId, id della visita e risposta all'invito
* @returns {int}
 */
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
