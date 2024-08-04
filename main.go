package main

import (
	"log"

	"github.com/devmor-j/basic-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func healthcheckHandler(c *fiber.Ctx) error {
	return c.SendString("ok")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheckHandler)

	app.Get("/products", handlers.GetAllProducts)
	app.Post("/products", handlers.CreateProduct)

	log.Fatal(app.Listen(":3000"))

}
