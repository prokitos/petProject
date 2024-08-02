package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"module/internal/config"
	"module/internal/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func sendToAuth(c *fiber.Ctx, sendData models.TokenResponser, supAddress string) (models.Tokens, error) {

	baseURL, _ := url.Parse(config.ExternalAddress.AuthService + supAddress)

	params := url.Values{}
	params.Add("login", sendData.Login)
	params.Add("password", sendData.Password)
	params.Add("password_confirm", sendData.PasswordConfirm)
	baseURL.RawQuery = params.Encode()

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.PostForm(baseURL.String(), params)
	if err != nil {
		log.Debug("no connect to auth service")
		return models.Tokens{}, models.ResponseConnectError()
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var instance models.ExternalStruct
	json.Unmarshal(body, &instance)

	return instance.Data, errors.New(instance.Status)

}

func sendTokenToCheck(c *fiber.Ctx, token string, supAddress string) (int, error) {

	baseURL := config.ExternalAddress.AuthService + supAddress

	req, err := http.NewRequest("POST", baseURL, nil)
	req.Header.Add("Authorization", token)

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Debug("no connect to auth service")
		return 0, models.ResponseConnectError()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("reading byte error")
		return 0, models.ResponseEncodingError()
	}

	var instance models.ExternalStruct
	json.Unmarshal(body, &instance)

	return instance.Access, errors.New(instance.Status)

}

func sendTokenToRefresh(c *fiber.Ctx, token string, refresh string, supAddress string) models.ExternalStruct {

	baseURL, _ := url.Parse(config.ExternalAddress.AuthService + supAddress)
	params := url.Values{}
	params.Add("refresh", refresh)
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", baseURL.String(), nil)
	req.Header.Add("Authorization", token)

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Debug("no connect to auth service")
		return models.ExternalStruct{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("reading byte error")
		return models.ExternalStruct{}
	}

	var instance models.ExternalStruct
	json.Unmarshal(body, &instance)

	return instance

}
