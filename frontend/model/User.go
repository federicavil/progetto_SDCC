package model

type Credential struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Login struct {
	Credential           Credential
	PasswordConfirmation string `form:"confirmPassword"`
}
