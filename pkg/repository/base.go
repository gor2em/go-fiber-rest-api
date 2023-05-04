package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Get(id uint, model interface{}) error
	Create(model interface{}) error
	Update(id uint, model interface{}) error
	Delete(id uint, model interface{}) error
	Find(model interface{}, where ...interface{}) error
}

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db: db}
}

func (repo *baseRepository) Get(id uint, model interface{}) error {
	err := repo.db.First(id, model).Error

	if err != nil {
		return err
	}
	
	return nil

}

func (repo *baseRepository) Create(model interface{}) error {
	err := repo.db.Create(model).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *baseRepository) Update (id uint,model interface{}) error{
	err := repo.db.Save(model).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *baseRepository) Delete (id uint, model interface{}) error{
	err := repo.db.Delete(id, model).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *baseRepository) Find(model interface{}, where ...interface{}) error {
	err := repo.db.Find(model, where...).Error;
	
	if err != nil{
		return err;
	}

	return nil
}