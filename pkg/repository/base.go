package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Get(id uint, model interface{}) error
	Create(model interface{}) error
	Update(id uint, model interface{}) error
	Delete(id uint, model interface{}) error
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

func (repo *baseRepository) Create(model interface{}) error {
	return repo.db.Create(model).Error
}

func (repo *baseRepository) Update (id uint,model interface{}) error{
	return repo.db.Save(model).Error
}

func (repo *baseRepository) Delete (id uint, model interface{}) error{
	return repo.db.Delete(id, model).Error
}