package service

import (
	"errors"
	"go-fiber-rest-api/pkg/model"
	"go-fiber-rest-api/pkg/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	err := s.userRepo.Get(id, user)
	if err != nil {
		return nil, errors.New("failed to get user by ID")
	}
	return user, nil
}