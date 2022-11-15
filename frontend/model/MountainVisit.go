package model

type MountainVisit struct {
	Pathname  string `db:"Pathname" json:"Pathname" form:"name"`
	Username  string `db:"Username" json:"Username"`
	Timestamp string `db:"Timestamp" json:"Timestamp" form:"visitTime"`
}
