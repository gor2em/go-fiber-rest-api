package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() (*gorm.DB, error) {

    // loadEnv()

    // host := os.Getenv("DB_HOST")
    // username := os.Getenv("DB_USER")
    // password := os.Getenv("DB_PASSWORD")
    // databaseName := os.Getenv("DB_NAME")
    // port := os.Getenv("DB_PORT")

    host := "localhost"
    username := "root"
    password := "secret"
    databaseName := "test"
    port := "5432"

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// func loadEnv(){
// 	err := godotenv.Load(".env.local")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

