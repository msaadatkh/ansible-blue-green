package main

import (
	"time"

	"github.com/gofiber/fiber"
)

func main() {

	healthy := false

	time.AfterFunc(time.Second*10, func() {
		healthy = true
	})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Welcome to my awesome app v3.0.0!")
	})
	app.Get("/healthz", func(c *fiber.Ctx) {
		if healthy {
			c.Status(200)
		} else {
			c.Status(500)
		}
	})

	app.Listen(3000)
}
