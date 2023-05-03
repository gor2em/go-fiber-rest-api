package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Get(id uint, model interface{}) error
}

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db: db}
}

func (repo *baseRepository) Get(id uint, model interface{}) error {
	return repo.db.First(model, id).Error
}

