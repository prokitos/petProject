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

type ResponseStr struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
	Cars        []Car  `json:"cars"               example:"...."`
}
