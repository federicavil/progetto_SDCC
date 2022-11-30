package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"sync"
	"time"
)

type NotificationController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: imposta la view da mostrare all'utente
 */
func (this *NotificationController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "notifications.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login.
* Altrimenti, riceve informazioni sulle notifiche dalla sessione e le mostra sulla view
 */
func (this *NotificationController) Get() {
	isLogged := CheckLogin(this.session, "/notifications")
	if !isLogged {
		err := this.session.Set("prevPage", "notifications")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else {
		notifications := this.session.Get("notifications")
		if notifications != nil {
			this.Data["notifications"] = notifications.([]model.Notification)
		}
	}
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di effettuare
* l'operazione di accept/decline di una certo invito
 */
func (this *NotificationController) Post() {
	accept := this.GetString("Accept")
	decline := this.GetString("Decline")
	viewInfo := this.GetString("ViewInfo")
	var visitId string
	if accept != "" || decline != "" {
		var response string
		// Chiamo api_gateway
		if accept != "" {
			response = "true"
			visitId = accept
		} else {
			response = "false"
			visitId = decline
		}
		req := httplib.Post("http://" + conf.GetApiGateway() + "/acceptOrRefuseInvite")
		req.Param("response", response)
		req.Param("id_visit", visitId)
		req.Param("username", this.session.Get("userId").(string))
		str, _ := req.Bytes()
		fmt.Println(str)
		// Cancello la notifica dalla lista nel frontend
		mutex := this.session.Get("sessionMutex").(*sync.Mutex)
		mutex.Lock()
		notifications := this.session.Get("notifications").([]model.Notification)
		i := 0
		for i = 0; i < len(notifications); i++ {
			if notifications[i].IdVisit == accept || notifications[i].IdVisit == decline {
				break
			}
		}
		if i == len(notifications)-1 {
			if i == 0 {
				notifications = nil
			} else {
				notifications = append(notifications[:i])
			}
		} else {
			notifications = append(notifications[:i], notifications[i+1:]...)
		}
		this.session.Set("notification", notifications)
		mutex.Unlock()

	} else if viewInfo != "" {
		// Redirect su view visit info
		req := httplib.Get("http://" + conf.GetApiGateway() + "/getVisitByID")
		req.Param("visitId", viewInfo)
		str, _ := req.Bytes()
		visit := model.MountainVisitComplete{}
		err := json.Unmarshal(str, &visit)
		if err != nil {
			return
		}
		err = this.session.Set("selectedVisit", visit)
		if err != nil {
			return
		}
		print("\nSONO IN NOTIFICATION E HO: ")
		fmt.Println(this.session.Get("selectedVisit"))
		this.Redirect("viewVisitInfo", 302)
		print("HO FATTO REDIRECT")
	}
}

/*
* Funzione eseguita come goroutine per ricevere notifiche relative agli inviti a visite da parte di altri utenti.
* Ogni 10 secondi si comunica con l'API Gateway per ricevere notifiche e mantenerle nella sessione.
* @param {string}: userId dell'utente ricevente (appena loggato) di cui si vogliono recuperare gli inviti
* @param {session.Store}: identificatore della sessione, usata per mantenere info sugli inviti ricevuti
 */
func NotificationPolling(userId string, session session.Store) {
	for true {
		if session.Get("userId") == "" {
			return
		}
		// chiama api_gateway per vedere se ci sono nuove notifiche
		req := httplib.Get("http://" + conf.GetApiGateway() + "/notifications")
		req.Param("userId", userId)
		req.Param("isLogged", "true")
		notificationsjson, _ := req.Bytes()
		var notificationString string
		_ = json.Unmarshal(notificationsjson, &notificationString)
		var notifications []model.Notification
		_ = json.Unmarshal([]byte(notificationString), &notifications)
		// se ci sono le mette sulla sessione
		/*var notifications []model.Notification
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
		fmt.Println("Sono il thread: sto per andare in sleep")*/
		mutex := session.Get("sessionMutex").(*sync.Mutex)
		mutex.Lock()
		/*		notifications = session.Get("notification").([]model.Notification)
				notifications = append(notifications, model.Notification{
					IdVisit: 2,
					Visit: model.MountainVisit{
						Pathname:  "Gran Sasso",
						Username:  "diana",
						Timestamp: "2459902.2131944443",
					},
					Participants: participants,
				})*/
		if session.Get("notifications") != nil {
			if len(notifications) != 0 {
				session.Set("notifications", notifications)
			}
		} else {
			session.Set("notifications", notifications)
		}
		mutex.Unlock()
		time.Sleep(10 * time.Second)
	}
}
