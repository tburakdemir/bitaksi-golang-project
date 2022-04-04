package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tburakdemir/driver-location-api/configs"
	"github.com/tburakdemir/driver-location-api/models"
	"github.com/tburakdemir/driver-location-api/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var locationCollection *mongo.Collection = configs.GetCollection(configs.DB, "locations")
var validate = validator.New()

func CreateLocation(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var location models.Location
    defer cancel()

    //validate the request body
    if err := c.BodyParser(&location); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.LocationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
	fmt.Println(location)

    //use the validator library to validate required fields
    if validationErr := validate.Struct(&location); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.LocationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

    newLocation := models.Location{
		Type: 	  location.Type,
		Coordinates: location.Coordinates,
    }

	fmt.Println("new loc: ", newLocation)

    result, err := locationCollection.InsertOne(ctx, bson.M{"location": newLocation})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.LocationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    return c.Status(http.StatusCreated).JSON(responses.LocationResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}