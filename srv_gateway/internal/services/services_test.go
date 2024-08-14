package services

import (
	"module/internal/models"
	"testing"

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
			expected:    models.ResponseBadRequest().Error(),
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

		resp := Authorization(c)
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
			expected:    models.ResponseBadRequest().Error(),
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

		resp := Register(c)
		assert.Equalf(t, test.expected, resp.Error(), test.description)
	}
}
