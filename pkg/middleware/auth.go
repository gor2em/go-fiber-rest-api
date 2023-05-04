package middleware

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Auth(c *fiber.Ctx) error {

	//get token
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}


	//verify token
	token = strings.Replace(token, "Bearer ", "", 1)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY_XYZ123"), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	userId, err := uuid.Parse(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//set token user id
	c.Locals("userId", userId)

	return c.Next()
}
