package model

import (
	"time"

	"github.com/google/uuid"
)

// User Model
type User struct {
	ID       	*uuid.UUID   	`gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name	 	string 			`gorm:"type:varchar(50);default:null"`
	Surname	 	string 			`gorm:"type:varchar(50);default:null"`
	Email    	string 			`gorm:"type:varchar(100);uniqueIndex;not null"`
	Username    string 			`gorm:"type:varchar(30);uniqueIndex;not null"`
	Password    string 			`gorm:"type:varchar(60);not null"`
	Company    	string 			`gorm:"type:varchar(50);not null"`
	CreatedAt 	*time.Time 		`gorm:"not null;default:now()"`
	UpdatedAt 	*time.Time 		`gorm:"not null;default:now()"`
}


type RegisterRequest struct{
	Name		string			`validate:"required,min=3"`
	Surname		string			`validate:"required,min=3"`
	Email		string			`validate:"required,email"`
	Username	string			`validate:"required,min=3,max=60"`
	Password	string			`validate:"required,min=3"`
	Company		string			`validate:"required,min=3"`
}

type LoginRequest struct {
	Email    	string 			`validate:"required,email"`
	Password 	string 			`validate:"required,min=3,max=60"`
}

type UserResponse struct{
	ID        	uuid.UUID 		`json:"id,omitempty"`
	Name		string			`json:"name,omitempty"`
	Surname		string			`json:"surname,omitempty"`
	Email		string			`json:"email"`
	Username	string			`json:"username"`
	Company		string			`json:"company"`
	CreatedAt 	time.Time 		`json:"created_at"`
}


func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        		*user.ID,
		Name: 			user.Name,
		Surname: 		user.Surname,
		Username:      	user.Username,
		Email:     		user.Email,
		CreatedAt: 		*user.CreatedAt,
		Company: 		user.Company,
	}
}