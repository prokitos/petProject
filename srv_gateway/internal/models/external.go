package models

type TokenResponser struct {
	Login           string
	Password        string
	PasswordConfirm string
}

type RequestParse struct {
	Status string
}
