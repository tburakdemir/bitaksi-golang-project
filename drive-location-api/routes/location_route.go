package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tburakdemir/driver-location-api/controllers"
)

func LocationRoute(app *fiber.App) {
     app.Post("/locations", controllers.CreateLocation)
	 app.Get("/locations", controllers.GetLocations)
}