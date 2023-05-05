package handler

import (
	"go-fiber-rest-api/pkg/middleware"
	"go-fiber-rest-api/pkg/model"
	"go-fiber-rest-api/pkg/service"
	"go-fiber-rest-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}


func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	// check body
	if err := c.BodyParser(&req); err != nil {
		return middleware.ErrorHandler(c, err, &fiber.ErrBadRequest.Code, &fiber.ErrBadRequest.Message)
	}

	// check req validations
	errors := middleware.ValidateStruct(req)
    if errors != nil {
       return c.Status(fiber.StatusBadRequest).JSON(errors)
    }

	//service
	user, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok{
			return middleware.ErrorHandler(c, err, &e.Code, &e.Message)
		}
	}

	//create token
	token, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
			errorMessage := "Failed to generate JWT token"
			return middleware.ErrorHandler(c, err, nil, &errorMessage)
	}

	//send user
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    model.FilterUserRecord(user),
		"token":token,
	})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return middleware.ErrorHandler(c, err, &fiber.ErrBadRequest.Code, &fiber.ErrBadRequest.Message)
	}
	
	errors := middleware.ValidateStruct(req)
	if errors != nil {
	   return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	
	hashedPassword, _ := utils.GenerateHashPassword(req.Password)

	user := &model.User{
		Name: req.Name,
		Surname: req.Surname,
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
		Company: req.Company,
	}

	user, err := h.userService.Register(user)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok{
			return middleware.ErrorHandler(c, err, &e.Code, &e.Message)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    model.FilterUserRecord(user),
	})
}