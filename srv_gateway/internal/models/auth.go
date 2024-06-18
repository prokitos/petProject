package models

type ExternalStruct struct {
	Status string
	Data   Tokens
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
