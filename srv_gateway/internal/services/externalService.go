package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"module/internal/config"
	"module/internal/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
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
		fmt.Println("no connect to auth service")
		return models.Tokens{}, errors.New("not connection to external service")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	models := convertResult(body)
	return models.Data, errors.New(models.Status)

}

func convertResult(res []byte) models.ExternalStruct {
	var instance models.ExternalStruct
	json.Unmarshal(res, &instance)

	return instance
}
