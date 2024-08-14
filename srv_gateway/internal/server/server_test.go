package server

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHttpConnect(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
		method       string
	}{
		{
			description:  "get HTTP status 201, login",
			route:        "/login",
			expectedCode: 201,
			method:       "POST",
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
			method:       "GET",
		},
		{
			description:  "get HTTP status 201, register",
			route:        "/register",
			expectedCode: 201,
			method:       "POST",
		},
		{
			description:  "get HTTP status 201, refresh",
			route:        "/refresh",
			expectedCode: 201,
			method:       "POST",
		},
		{
			description:  "get HTTP status 201, debug",
			route:        "/debugBoost",
			expectedCode: 201,
			method:       "POST",
		},
		{
			description:  "get HTTP status 201, car insert",
			route:        "/car",
			expectedCode: 500,
			method:       "POST",
		},
		{
			description:  "get HTTP status 201, car delete",
			route:        "/car",
			expectedCode: 500,
			method:       "DELETE",
		},
		{
			description:  "get HTTP status 201, car show",
			route:        "/car",
			expectedCode: 500,
			method:       "GET",
		},
		{
			description:  "get HTTP status 201, car update",
			route:        "/car",
			expectedCode: 500,
			method:       "PUT",
		},
		{
			description:  "get HTTP status 404, selling insert",
			route:        "/carSell",
			expectedCode: 500,
			method:       "POST",
		},
		{
			description:  "get HTTP status 201, selling delete",
			route:        "/carSell",
			expectedCode: 500,
			method:       "DELETE",
		},
		{
			description:  "get HTTP status 201, selling show",
			route:        "/carSell",
			expectedCode: 500,
			method:       "GET",
		},
		{
			description:  "get HTTP status 201, selling update",
			route:        "/carSell",
			expectedCode: 500,
			method:       "PUT",
		},
	}

	app := fiber.New()
	handlers(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
