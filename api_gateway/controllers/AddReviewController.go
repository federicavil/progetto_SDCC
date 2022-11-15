package controllers

import (
	"api_gateway/proto"
	"context"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AddReviewController struct {
	web.Controller
}

func (this *AddReviewController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

func (this *AddReviewController) Post() {
	review := this.Ctx.Input.Query("review")
	fmt.Println("aggiunta review" + review)
	conn, _ := Dial("9091")
	client := proto.NewReviewManagerServiceClient(conn)
	response, err := client.AddReview(context.TODO(), &proto.AddReviewRequest{Review: review})
	if err != nil {
		fmt.Println(err)
	}
	this.Ctx.WriteString(response.AddedReview)
}
