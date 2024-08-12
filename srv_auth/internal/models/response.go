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

func ResponseServerConnectionError() error {
	return errors.New("connection to server error")
}

func ErrorLoginPassword() error {
	return errors.New("login or password error")
}

func ResponseAccessError() error {
	return errors.New("access error")
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
func ResponseTokenDontSame() error {
	return errors.New("token dont same")
}
func ResponseTokenUnauthorized() error {
	return errors.New("unauthorized token")
}
func ResponseTokenDontRead() error {
	return errors.New("token hasnt deconvert")
}

func ResponseUserExist() error {
	return errors.New("user already exist")
}
