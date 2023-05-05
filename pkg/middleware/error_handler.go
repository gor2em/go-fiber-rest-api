package middleware

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error, statusCode *int, message *string) error {

	if statusCode == nil {
		statusCode = &fiber.ErrInternalServerError.Code
	}

	if message == nil{
		message = &fiber.ErrInternalServerError.Message
	}

    if e, ok := err.(*fiber.Error); ok {
        *statusCode = e.Code
        *message = e.Message
    }

    return c.Status(*statusCode).JSON(fiber.Map{
        "status": *statusCode,
        "message": *message,
    })
}