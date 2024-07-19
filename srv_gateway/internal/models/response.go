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

type ResponseCar struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
	Cars        []Car  `json:"cars"               example:"...."`
}

func ResponseCarGoodCreate() ResponseCar {
	var resp ResponseCar
	resp.Code = 200
	resp.Description = "car create success"
	resp.Cars = nil
	return resp
}
func ResponseCarBadCreate() ResponseCar {
	var resp ResponseCar
	resp.Code = 400
	resp.Description = "car create failed"
	resp.Cars = nil
	return resp
}

func ResponseCarBadShow() ResponseCar {
	var resp ResponseCar
	resp.Code = 400
	resp.Description = "car get failed"
	resp.Cars = nil
	return resp
}
func ResponseCarGoodShow(cars []Car) ResponseCar {
	var resp ResponseCar
	resp.Code = 200
	resp.Description = "car get success"
	resp.Cars = cars
	return resp
}

func ResponseCarUnsupported() ResponseCar {
	var resp ResponseCar
	resp.Code = 404
	resp.Description = "unexpected error"
	resp.Cars = nil
	return resp
}

func ResponseCarBadDelete() ResponseCar {
	var resp ResponseCar
	resp.Code = 400
	resp.Description = "nothing to delete"
	resp.Cars = nil
	return resp
}

func ResponseCarGoodDelete() ResponseCar {
	var resp ResponseCar
	resp.Code = 200
	resp.Description = "delete success"
	resp.Cars = nil
	return resp
}

func ResponseCarBadUpdate() ResponseCar {
	var resp ResponseCar
	resp.Code = 400
	resp.Description = "error at update"
	resp.Cars = nil
	return resp
}

func ResponseCarGoodUpdate() ResponseCar {
	var resp ResponseCar
	resp.Code = 200
	resp.Description = "success update"
	resp.Cars = nil
	return resp
}

type ResponseSell struct {
	Description string    `json:"description"        example:"description"`
	Code        int       `json:"code"               example:"status"`
	Sells       []Selling `json:"cars"               example:"...."`
}

func ResponseSellGoodShow(sells []Selling) ResponseSell {
	var resp ResponseSell
	resp.Code = 200
	resp.Description = "car get success"
	resp.Sells = sells
	return resp
}
func ResponseSellBadShow() ResponseSell {
	var resp ResponseSell
	resp.Code = 400
	resp.Description = "car get failed"
	resp.Sells = nil
	return resp
}

func ResponseSellGoodExecute() ResponseSell {
	var resp ResponseSell
	resp.Code = 200
	resp.Description = "selling operation success"
	resp.Sells = nil
	return resp
}
func ResponseSellBadExecute() ResponseSell {
	var resp ResponseSell
	resp.Code = 400
	resp.Description = "selling operation failed"
	resp.Sells = nil
	return resp
}

func ResponseTokenGood() error {
	return errors.New("token is useful")
}

func ResponseConnectError() error {
	return errors.New("not connection to external services")
}

func ResponseEncodingError() error {
	return errors.New("error encoding response of server")
}

func ResponseTokenExpired() error {
	return errors.New("token is expired")
}
