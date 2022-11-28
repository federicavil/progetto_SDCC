package controllers

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type ViewReviewsController struct {
	web.Controller
}

/*
* Gestione chiamata GET: effettua chiamata grpc per recuperare la lista delle recensioni di un dato percorso
* @returns {string}: lista delle recensioni in json
 */
func (this *ViewReviewsController) Get() {
	mountainPathName := this.Ctx.Input.Query("mountainPathName")
	conn, _ := Dial(conf.GetConnectionConf("review_manager"))
	client := proto.NewReviewManagerServiceClient(conn)
	response, e := client.GetReviews(context.TODO(), &proto.GetReviewsRequest{MountainPathName: mountainPathName})
	if e != nil {
		fmt.Println(e)
	}
	this.Ctx.WriteString(response.ReviewsList)

}
