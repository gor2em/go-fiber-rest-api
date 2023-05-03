package main

import (
	"fmt"
	"go-fiber-rest-api/pkg/config"
	"go-fiber-rest-api/pkg/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database, err := config.ConnectDB()
	
	if err != nil{
		log.Fatalf("cannot connect to db %v", err)
	}

	fmt.Println(database.Config)

	app := fiber.New()

	//users
	routes.UserRoutes(app,database)

	app.Listen(":8000")

}