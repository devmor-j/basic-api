package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func healthcheckHandler(c *fiber.Ctx) error {
	return c.SendString("ok")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheckHandler)

	log.Fatal(app.Listen(":3000"))
}
