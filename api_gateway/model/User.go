package model

type Credential struct {
	Username string
	Password string
}

type Login struct {
	Credential           Credential
	PasswordConfirmation string
}
