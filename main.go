package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	err := app.Listen("localhost:8124")

	if err != nil {
		panic(err)
	}
}
