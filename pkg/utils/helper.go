package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func GenerateHashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func CompareHashPassword(hashedPassword,password string ) error{
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil	
}

func GenerateAccessToken(userID *uuid.UUID, userEmail string) (string, error) {
	// Define JWT claims
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["email"] = userEmail
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // Token expires in 7 days

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//from env
	return token.SignedString([]byte("SECRET_KEY_XYZ123"))
}
