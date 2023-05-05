package repository

import (
	"go-fiber-rest-api/pkg/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository

	FindUserByEmail(email string) (*model.User, error)
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

func (repo *userRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
