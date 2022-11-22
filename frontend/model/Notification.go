package model

type Notification struct {
	IdVisit      string   `json:"visit"`
	Username     string   `json:"creator"`
	Participants []string `json:"participants"`
}
