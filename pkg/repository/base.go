package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Get(id uint, model interface{}) error
	Create(model interface{}) (interface{}, error)
	Update(id uint, model interface{}) error
	Delete(id uint, model interface{}) error
	Where(model interface{}, query string, args... interface{}) (bool, error)
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

func (repo *baseRepository) Create(model interface{}) (interface{}, error) {
	err := repo.db.Create(model).Error
	return model, err
}

func (repo *baseRepository) Update(id uint,model interface{}) error{
	err := repo.db.Save(model).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *baseRepository) Delete(id uint, model interface{}) error{
	err := repo.db.Delete(id, model).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *baseRepository) Where(model interface{}, query string, args ...interface{}) (bool, error) {
	var count int64
	err := repo.db.Model(model).Where(query, args...).Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}