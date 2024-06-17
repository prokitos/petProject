package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"module/internal/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

var secretKey string = "gpt45"

func sendToSecond(c *fiber.Ctx, sendData models.TokenResponser, supAddress string) {

	baseURL, _ := url.Parse("http://localhost:8002/" + supAddress)

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
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	convertResult(body)

}

func convertResult(res []byte) {
	var test TestStruct
	json.Unmarshal(res, &test)

	fmt.Println(test.Data)
	fmt.Println()
	fmt.Println(test)
}

type TestStruct struct {
	Status error
	Data   TokenResponser
}

type TokenResponser struct {
	AccessToken  string
	RefreshToken string
}
