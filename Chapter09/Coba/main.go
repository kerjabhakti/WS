package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	app.Listen(":3000")
}
