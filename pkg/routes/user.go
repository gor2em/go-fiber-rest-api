package routes

import (
	"go-fiber-rest-api/pkg/api/handler"
	"go-fiber-rest-api/pkg/repository"
	"go-fiber-rest-api/pkg/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(router *fiber.App, db *gorm.DB){
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	
	userRouter := router.Group("/users")	
	userRouter.Post("/register", userHandler.Register)
	userRouter.Post("/login", userHandler.Login)
	
	//update
	//forgot-password
	//reset-password
}