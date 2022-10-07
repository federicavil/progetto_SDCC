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

type SimpleSearchStruct struct {
	Name string
}

type AdvancedSearchStruct struct {
	City       string
	Province   string
	Region     string
	Level      string
	Cyclable   int
	Family     int
	Historical int
}

const (
	T   string = "T"
	E   string = "E"
	EE  string = "EE"
	EEA string = "EEA"
)
