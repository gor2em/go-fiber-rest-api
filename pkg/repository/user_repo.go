package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository
}

type userRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository(db),
		db:db,
	}
}