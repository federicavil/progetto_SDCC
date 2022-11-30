package model

type Review struct {
	MountainPathName string `json:"mountainPathName"`
	Vote             int    `json:"vote"`
	Title            string `json:"title"`
	Comment          string `json:"comment"`
	Author           string `json:"author"`
}

type ReviewForm struct {
	Vote    int    `form:"vote"`
	Title   string `form:"title"`
	Comment string `form:"comment"`
}
