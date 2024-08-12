package services

import (
	"module/internal/models"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestAuth(t *testing.T) {

	tests := []struct {
		description string
		expected    string
		paramName   []string
		paramValue  []string
	}{
		{
			description: "auth login and password error",
			expected:    models.ResponseBadRequest().Error(),
			paramName:   nil,
			paramValue:  nil,
		},
		{
			description: "auth connection with bd error",
			expected:    models.ResponseServerConnectionError().Error(),
			paramName:   []string{"login", "password"},
			paramValue:  []string{"profiles", "1234567"},
		},
	}

	for _, test := range tests {
		app := fiber.New()
		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		for i := 0; i < len(test.paramName); i++ {
			c.Context().QueryArgs().Add(test.paramName[i], test.paramValue[i])
		}

		_, resp := Authorization(c)
		assert.Equalf(t, test.expected, resp.Error(), test.description)
	}
}

func TestRegistration(t *testing.T) {

	tests := []struct {
		description string
		expected    string
		paramName   []string
		paramValue  []string
	}{
		{
			description: "registration login and password error",
			expected:    models.ResponseBadRequest().Error(),
			paramName:   []string{"login", "password"},
			paramValue:  []string{"profi", "12345"},
		},
		{
			description: "registrion connection with bd error",
			expected:    models.ResponseServerConnectionError().Error(),
			paramName:   []string{"login", "password", "password_confirm"},
			paramValue:  []string{"profiles", "1234567", "1234567"},
		},
	}

	for _, test := range tests {
		app := fiber.New()
		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		for i := 0; i < len(test.paramName); i++ {
			c.Context().QueryArgs().Add(test.paramName[i], test.paramValue[i])
		}

		_, resp := Registration(c)
		assert.Equalf(t, test.expected, resp.Error(), test.description)
	}
}

func TestTokenCreate(t *testing.T) {

	tests := []struct {
		description string
		notexpected string
		user        models.Users
	}{
		{
			description: "token create with epty parameters",
			notexpected: "",
			user:        models.Users{},
		},
		{
			description: "normal token create",
			notexpected: "",
			user:        models.Users{Login: "profile", Password: "1234567", AccessLevel: 3, RefreshToken: "refresh"},
		},
	}

	for _, test := range tests {

		resp := createTokenAccess(test.user)
		assert.NotEqualf(t, test.notexpected, resp, test.description)
	}

}

func TestValidateToken(t *testing.T) {

	AccessDuration = time.Minute * 10

	tests := []struct {
		description string
		expected    string
		token       string
	}{
		{
			description: "normal token check",
			expected:    models.ResponseTokenGood().Error(),
			token:       "bearer " + createTokenAccess(models.Users{Login: "profile", Password: "1234567", AccessLevel: 3, RefreshToken: "refresh"}),
		},
		{
			description: "old token check",
			expected:    models.ResponseTokenExpired().Error(),
			token:       "bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InBsYXllciIsIkFjY2Vzc0xldmVsIjo0LCJleHAiOjE3MjE2NTExMjZ9.zY9BUyGUIF6d4S_pcseADvb9leYFKI_uohqTq-kLDNPSOA173oO6NAydDC6cNRBusDEwkWF2UXKTTC8vsqqTtQ",
		},
		{
			description: "error token check",
			expected:    models.ResponseTokenError().Error(),
			token:       "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpAiOjE3MjE2NTExMjZ9.zY9BUyGUIFNPSOA173oO6NAydDC6cNRBusDEwkWF2UXKTTC8vsqqTtQ",
		},
	}

	for _, test := range tests {

		resp, _ := TokenAccessValidate(test.token)
		assert.Equalf(t, test.expected, resp, test.description)
	}
}
