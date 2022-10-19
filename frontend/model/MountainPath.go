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

type MountainPathRet struct {
	Name       string `db:"Name" json:"Name"`
	Altitude   int    `db:"Altitude" json:"Altitude"`
	City       string `db:"City" json:"City"`
	Province   string `db:"Province" json:"Province"`
	Region     string `db:"Region" json:"Region"`
	Length     int    `db:"Length" json:"Length"`
	Level      string `db:"Level" json:"Level"`
	Cyclable   string `db:"Cyclable" json:"Cyclable"`
	Family     string `db:"Family" json:"Family"`
	Historical string `db:"Historical" json:"Historical"`
}

type MountainPathInterface interface {
	Name() string
	Altitude() int
	City() string
	Region() string
	Province() string
	Length() int
	Level() string
	Cyclable() bool
	Family() bool
	Historical() bool
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
