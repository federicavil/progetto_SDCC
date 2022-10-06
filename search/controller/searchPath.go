package controller

import (
	"fmt"
	"search/model"
)

type Args struct {
	Name string
}

type Search int

func (t *Search) SimpleSearch(args *Args, reply *[]model.MountainPath) error {
	fmt.Println("AO")
	fmt.Println(*args)
	path_1 := model.MountainPath{
		Name:     "Gran Sasso",
		Altitude: 1234,

		City:     "pippo",
		Province: "Aquila",
		Region:   "Abruzzo",

		Length:     1094,
		Level:      "EE",
		Cyclable:   false,
		Family:     false,
		Historical: false,
	}
	path_2 := model.MountainPath{
		Name:       "Piccolo Sasso",
		Altitude:   1234,
		City:       "pippo",
		Province:   "Aquila",
		Region:     "Abruzzo",
		Length:     1094,
		Level:      "EE",
		Cyclable:   false,
		Family:     true,
		Historical: false,
	}
	*reply = []model.MountainPath{path_1, path_2}
	return nil
}
