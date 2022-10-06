package model

type MountainPath struct {
	Name       string
	Altitude   int
	City       string
	Region     string
	Province   string
	Length     int
	Level      string
	Cyclable   bool
	Family     bool
	Historical bool
}

const (
	T   string = "T"
	E   string = "E"
	EE  string = "EE"
	EEA string = "EEA"
)
