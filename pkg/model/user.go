package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       	*uuid.UUID   	`gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name	 	string 			`gorm:"type:varchar(50)"`
	Surname	 	string 			`gorm:"type:varchar(50)"`
	Email    	string 			`gorm:"type:varchar(100);uniqueIndex;not null"`
	Username    string 			`gorm:"type:varchar(30);uniqueIndex;not null"`
	Password    string 			`gorm:"type:varchar(50);not null"`
	CreatedAt 	*time.Time 		`gorm:"not null;default:now()"`
	UpdatedAt 	*time.Time 		`gorm:"not null;default:now()"`
}

type SignUpInput struct{
	Name		string			`json:"name"`
	Surname		string			`json:"surname"`
	Email		string			`json:"email" validate:"required, email"`
	Username	string			`json:"username" validate:"required, min=3"`
	Password	string			`json:"password" validate:"required, min=3"`
}