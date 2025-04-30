package service

import (
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/model"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/repository"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}
func (s *userServiceImpl) CreateUser(name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}

	return s.userRepo.CreateUser(user)
}
