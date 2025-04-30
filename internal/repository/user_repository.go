package repository

import "github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
}
