package models

import "github.com/dgrijalva/jwt-go"

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

type TokenResponser struct {
	AccessToken  string
	RefreshToken string
}
