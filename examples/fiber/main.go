package main

import (
	"github.com/alecsavvy/allgood"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(allgood.Middleware())
	app.Get("/", handleSuccess)
	app.Get("/error", handleError)

	app.Listen(":3000")
}

func handleError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
}

func handleSuccess(c *fiber.Ctx) error {
	return c.SendString("Success")
}
