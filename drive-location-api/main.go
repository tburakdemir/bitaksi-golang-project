package main

import (
	"github.com/tburakdemir/driver-location-api/configs"
	"github.com/tburakdemir/driver-location-api/routes"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
    app := fiber.New()
	
    routes.LoginRoute(app)
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

    app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
    routes.LocationRoute(app)
  
    app.Listen("127.0.0.1:" + configs.Port())
}