package handler

import (
	"go-fiber-rest-api/pkg/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	// Get user ID from URL parameter
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Call service to get user by ID
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user"})
	}

	// Return user as JSON response
	return ctx.JSON(user)
}

// func (h handler) Create (c *fiber.Ctx) error {
// 	model := Model{}
// 	err := c.BodyParser(&model)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Error: err.Error()})
// 	}

// 	_, err = h.service.Create(model)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Error: err.Error()})
// 	}

// 	return c.Status(201).JSON("OK")
// }