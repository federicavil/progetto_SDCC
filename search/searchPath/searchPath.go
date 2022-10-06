package searchPath

import "fmt"

type Location struct {
	City     string
	Province string
	Region   string
}

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

type Args struct {
	Name string
}

type Search int

func (t *Search) SimpleSearch(args *Args, reply *[]MountainPath) error {
	fmt.Println("AO")
	fmt.Println(*args)
	path_1 := MountainPath{
		Name:     "Gran Sasso",
		Altitude: 1234,
		Location: Location{
			City:     "pippo",
			Province: "Aquila",
			Region:   "Abruzzo",
		},
		Length:     1094,
		Level:      "EE",
		Cyclable:   false,
		Family:     false,
		Historical: false,
	}
	path_2 := MountainPath{
		Name:     "Piccolo Sasso",
		Altitude: 1234,
		Location: Location{
			City:     "pippo",
			Province: "Aquila",
			Region:   "Abruzzo",
		},
		Length:     1094,
		Level:      "EE",
		Cyclable:   false,
		Family:     true,
		Historical: false,
	}
	*reply = []MountainPath{path_1, path_2}
	return nil
}
