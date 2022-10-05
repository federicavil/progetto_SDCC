package model

type MountainPath struct {
	Name       string
	Altitude   int
	Location   Location
	Length     int
	Level      string
	Cyclable   bool
	Family     bool
	Historical bool
}
