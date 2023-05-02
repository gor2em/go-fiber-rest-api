package main

import "go-fiber-rest-api/pkg/config"

func main() {
	
	config.LoadEnv()

	config.ConnectDB()

}