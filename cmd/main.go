package main

import (
	"fmt"
	"go-fiber-rest-api/pkg/config"
	"go-fiber-rest-api/pkg/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database, err := config.ConnectDB()
	
	if err != nil{
		log.Fatalf("cannot connect to db %v", err)
	}

	fmt.Println(database.Config)

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil{
		log.Fatal(err)
	}

	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	app.Post("/users/", handler.Create)

	app.Listen(":8000")

}