package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"strconv"
	"sync"
	"time"
)

type NotificationController struct {
	beego.Controller
	session session.Store
}

func (this *NotificationController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "notifications.html"
}

func (this *NotificationController) Get() {
	isLogged := CheckLogin(this.session, "/notifications")
	if !isLogged {
		err := this.session.Set("prevPage", "notifications")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else {
		this.Data["notifications"] = this.session.Get("notifications").([]model.Notification)
	}
}

func (this *NotificationController) Post() {
	fmt.Println("POST")
	accept := this.GetString("Accept")
	decline := this.GetString("Decline")
	viewInfo := this.GetString("ViewInfo")
	if accept != "" || decline != "" {
		// Chiamo api_gateway
		// Cancello la notifica dalla lista
		mutex := this.session.Get("sessionMutex").(*sync.Mutex)
		mutex.Lock()
		notifications := this.session.Get("notifications").([]model.Notification)
		i := 0
		for i = 0; i < len(notifications); i++ {
			if strconv.Itoa(notifications[i].IdVisit) == accept || strconv.Itoa(notifications[i].IdVisit) == decline {
				break
			}
		}
		notifications = append(notifications[:i], notifications[i+1:]...)
		this.session.Set("notification", notifications)
		mutex.Unlock()

	} else if viewInfo != "" {
		// Redirect su view visit info
	}
}

func NotificationPolling(username string, session session.Store) {
	for true {
		fmt.Println("Sono il thread")
		// chiama api_gateway per vedere se ci sono nuove notifiche
		// se ci sono le mette sulla sessione
		var notifications []model.Notification
		var participants []model.Participant
		participants = append(participants, model.Participant{
			IdVisit:  1,
			Username: "mattia",
			Answer:   true,
		})
		notifications = append(notifications, model.Notification{
			IdVisit: 1,
			Visit: model.MountainVisit{
				Pathname:  "Gran Sasso",
				Username:  "federica",
				Timestamp: "2459902.2131944443",
			},
			Participants: participants,
		})
		session.Set("notifications", notifications)
		fmt.Println("Sono il thread: sto per andare in sleep")
		time.Sleep(30 * time.Second)
		mutex := session.Get("sessionMutex").(*sync.Mutex)
		mutex.Lock()
		notifications = session.Get("notification").([]model.Notification)
		notifications = append(notifications, model.Notification{
			IdVisit: 2,
			Visit: model.MountainVisit{
				Pathname:  "Gran Sasso",
				Username:  "diana",
				Timestamp: "2459902.2131944443",
			},
			Participants: participants,
		})
		session.Set("notifications", notifications)
		mutex.Unlock()
		time.Sleep(30 * time.Second)
	}
}
