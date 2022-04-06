package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tburakdemir/driver-location-api/controllers"
)

func LoginRoute(app *fiber.App) {
     app.Post("/login", controllers.Login)
}