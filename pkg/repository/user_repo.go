package repository

import (
	"errors"
	"go-fiber-rest-api/pkg/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository

	FindUser(email string) (*model.User, error)
}

type userRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository(db),
		db:             db,
	}
}


func (repo *userRepository) FindUser(email string) (*model.User, error) {
	var user model.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	
	return &user, err
}
