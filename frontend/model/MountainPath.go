package model

type MountainPath struct {
	Name       string `db:"Name" json:"Name" form:"name"`
	Altitude   int    `db:"Altitude" json:"Altitude" form:"altitude"`
	City       string `db:"City" json:"City" form:"city"`
	Province   string `db:"Province" json:"Province" form:"province"`
	Region     string `db:"Region" json:"Region" form:"region"`
	Length     int    `db:"Length" json:"Length" form:"lenght"`
	Level      string `db:"Level" json:"Level" form:"level"`
	Cyclable   bool   `db:"Cyclable" json:"Cyclable" form:"cycleble"`
	Family     bool   `db:"Family" json:"Family" form:"familySuitable"`
	Historical bool   `db:"Historical" json:"Historical" form:"historicalElements"`
}

type SimpleSearchStruct struct {
	Name string
}

type AssistedSearchStruct struct {
	City       string ` form:"city"`
	Province   string ` form:"province"`
	Region     string ` form:"region"`
	Level      string ` form:"level"`
	Cyclable   int    ` form:"cyclable"`
	Family     int    ` form:"familySuitable"`
	Historical int    ` form:"historicalElements"`
}

const (
	T   string = "T"
	E   string = "E"
	EE  string = "EE"
	EEA string = "EEA"
)
