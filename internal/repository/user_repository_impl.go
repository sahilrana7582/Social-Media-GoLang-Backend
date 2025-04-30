package repository

import (
	"database/sql"
	"log"

	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/model"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {

	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	return user, nil
}
