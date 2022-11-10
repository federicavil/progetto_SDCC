package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

type AddReviewController struct {
	beego.Controller
	session session.Store
}

func (this *AddReviewController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "addReview.html"
}

func (this *AddReviewController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://127.0.0.1:5000/addReview")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
	if !isLogged {
		err := this.session.Set("prevPage", "addReview")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
}

func (this *AddReviewController) Post() {
	if this.GetString("backInfo") != "" {
		fmt.Println("back")
		this.Redirect("viewInfo", 302)
	} else {
		fmt.Println("POST REVIEW")
		var vote int
		/*for i := 5; i >= 1; i-- {
			star := string(i)
			if this.GetString(star) != "" {
				vote = i
				break
			}
		}*/
		btn1 := this.GetString("1")
		btn2 := this.GetString("2")
		btn3 := this.GetString("3")
		btn4 := this.GetString("4")
		btn5 := this.GetString("5")
		if btn5 != "" {
			vote = 5
		}
		if btn4 != "" {
			vote = 4
		}
		if btn3 != "" {
			vote = 3
		}
		if btn2 != "" {
			vote = 2
		}
		if btn1 != "" {
			vote = 1
		}

		fmt.Println(vote)
		path := this.session.Get("selectedPath").(model.MountainPath)

		review := model.Review{}
		review.Vote = vote
		review.Title = this.GetString("title")
		review.Comment = this.GetString("comment")
		review.Author = this.session.Get("userId").(string)
		review.MountainPathName = path.Name

		reviewJson, _ := json.Marshal(review)
		fmt.Println(string(reviewJson))
		req := httplib.Post("http://127.0.0.1:5000/addReview")
		req.Param("review", string(reviewJson))
		str, _ := req.Bytes()
		fmt.Println(str)
	}

}
