package model

type Review struct {
	MountainPathName string `json:"mountainPathName"`
	Vote             int    `json:"vote"`
	Title            string `json:"title"`
	Comment          string `json:"comment"`
	Author           string `json:"author"`
}
