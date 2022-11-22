package model

type MountainVisit struct {
	Pathname  string `db:"Pathname" json:"Pathname"`
	Username  string `db:"Username" json:"Username"`
	Timestamp string `db:"Timestamp" json:"Timestamp"`
}

type MountainVisitComplete struct {
	ID_Visit     int      `db:"ID_Visit" json:"ID_Visit"`
	ID_Path      string   `db:"ID_Path" json:"ID_Path"`
	Creator      string   `db:"Creator" json:"Creator"`
	Timestamp    string   `db:"Timestamp" json:"Timestamp"`
	Participants []string `db:"Participants" json:"Participants"`
}

type Participant struct {
	IdVisit  int
	Username string
	Answer   bool
}
