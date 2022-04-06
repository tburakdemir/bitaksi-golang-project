package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tburakdemir/driver-location-api/models"
	"github.com/tburakdemir/driver-location-api/responses"
)

func Login(c *fiber.Ctx) error {

	var login models.Login
	
	 if err := c.BodyParser(&login); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.LocationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	// Only for test purposes
	// Throws Unauthorized error
	if login.User != "admin" || login.Pass != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "Admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}