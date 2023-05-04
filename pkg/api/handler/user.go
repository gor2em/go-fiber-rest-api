package handler

import (
	"go-fiber-rest-api/pkg/model"
	"go-fiber-rest-api/pkg/service"
	"go-fiber-rest-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

type Response struct{
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginData model.SignInInput

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := h.userService.Login(loginData.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	err = utils.CompareHashPassword(user.Password, loginData.Password)
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	// Generate Token
	token, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate JWT token",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Login success.",
		"user":    model.FilterUserRecord(user),
		"status":fiber.StatusOK,
		"token":token,
	})
}

func (h *UserHandler) Register (c *fiber.Ctx) error {

	var payload *model.SignUpInput

	//validations
	// Parse request body
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	hashedPassword, _ := utils.GenerateHashPassword(payload.Password)

	user := &model.User{
		Name: payload.Name,
		Surname: payload.Surname,
		Username: payload.Username,
		Email: payload.Email,
		Password: hashedPassword,
		Company: payload.Company,
	}
	
    if err := h.userService.Register(user); err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "message": "User created successfully",
		"user":model.FilterUserRecord(user),
    })
}