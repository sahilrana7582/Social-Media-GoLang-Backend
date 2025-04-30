package service

import "github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/model"

type UserService interface {
	CreateUser(name, email string) (*model.User, error)
}
