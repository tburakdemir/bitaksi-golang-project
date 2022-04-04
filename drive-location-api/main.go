package main

import (
	"github.com/tburakdemir/driver-location-api/configs"
	"github.com/tburakdemir/driver-location-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
	routes.LocationRoute(app)
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })
  
    app.Listen("127.0.0.1:" + configs.Port())
}