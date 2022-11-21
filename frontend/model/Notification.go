package model

type Notification struct {
	IdVisit      int
	Visit        MountainVisit
	Participants []Participant
}
