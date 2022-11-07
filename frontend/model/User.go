package model

type Credential struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Login struct {
	Credential           Credential
	PasswordConfirmation string `form:"confirmPassword"`
}
type UserProfile struct {
	Name        string `form:"name" json:"name"`
	Surname     string `form:"surname" json:"surname"`
	Description string `form:"description" json:"description"`
}
