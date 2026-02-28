package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestRoutingHello(t *testing.T) {
	app := fiber.New()

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
