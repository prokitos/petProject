package models

import "errors"

func ResponseGood() error {
	return errors.New("good")
}

func ResponseErrorAtServer() error {
	return errors.New("internal error")
}

func ResponseBadRequest() error {
	return errors.New("bad request")
}

func ResponseTokenError() error {
	return errors.New("token error")
}
func ResponseTokenGood() error {
	return errors.New("token is useful")
}
func ResponseTokenExpired() error {
	return errors.New("token is expired")
}
func ResponseTokenUnauthorized() error {
	return errors.New("unauthorized token")
}

func ResponseUserExist() error {
	return errors.New("user already exist")
}
