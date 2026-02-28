package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestRoutingHello(t *testing.T) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)

	assert.Equal(t, "Hello", string(bytes))
}

func TestGetQuery(t *testing.T) {
	app.Get("/hello", func(ctx fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})

	// Test send query
	request := httptest.NewRequest("GET", "/hello?name=Arza", nil)
	response, err := app.Test(request)

	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Arza", string(bytes))

	// Test send without query, should return fallback value from name query
	request = httptest.NewRequest("GET", "/hello", nil)
	response, err = app.Test(request)

	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)

	assert.Equal(t, "Hello Guest", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/test-http", func(ctx fiber.Ctx) error {
		firstName := ctx.Get("first_name")       // from header
		middleName := ctx.Cookies("middle_name") // cookies

		return ctx.SendString("Hello " + firstName + " " + middleName)
	})

	request := httptest.NewRequest("GET", "/test-http", nil)
	request.Header.Set("first_name", "Arzaqul")
	request.AddCookie(&http.Cookie{Name: "middle_name", Value: "Mughny Al Fawwaz"})
	response, err := app.Test(request)

	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello Arzaqul Mughny Al Fawwaz", string(bytes))
}
