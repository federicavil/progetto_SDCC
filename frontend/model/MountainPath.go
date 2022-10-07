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
	City       string ` form:"city"`
	Province   string ` form:"province"`
	Region     string ` form:"region"`
	Level      string ` form:"level"`
	Cyclable   int    ` form:"cycleble"`
	Family     int    ` form:"historicalElements"`
	Historical int    ` form:"familySuitable"`
}

const (
	T   string = "T"
	E   string = "E"
	EE  string = "EE"
	EEA string = "EEA"
)
