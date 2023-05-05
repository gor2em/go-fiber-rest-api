package service

import (
	"errors"
	"go-fiber-rest-api/pkg/model"
	"go-fiber-rest-api/pkg/repository"
	"go-fiber-rest-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *model.User) (*model.User, error) {

	exists, err := s.userRepo.Where(&model.User{}, "email = ? OR username = ?", user.Email, user.Username)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fiber.NewError(fiber.ErrConflict.Code, fiber.ErrConflict.Message)
	}

	_, err = s.userRepo.Create(user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	return user, nil
}

func (s *UserService) Login(email, password string) (*model.User, error){
	user, err := s.userRepo.FindUserByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, fiber.NewError(fiber.ErrNotFound.Code, fiber.ErrNotFound.Message)
		}
	}

	err = utils.CompareHashPassword(user.Password, password)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrNotFound.Code, fiber.ErrNotFound.Message)
	}

	return user, nil
}

