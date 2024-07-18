package models

import "github.com/dgrijalva/jwt-go"

type ExternalStruct struct {
	Status string
	Data   Tokens
	Access int
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AccessToken struct {
	Login       string
	AccessLevel int
	jwt.StandardClaims
}

type RefreshToken struct {
	Login        string
	AccessLevel  int
	AcceessToken string
	jwt.StandardClaims
}
