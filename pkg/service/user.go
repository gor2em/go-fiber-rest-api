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

func (s *UserService) Register(user *model.User) error{
	return s.userRepo.Create(user)
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

