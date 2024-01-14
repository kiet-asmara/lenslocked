package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uint
	Email         string
	Password_hash string
}

type UserService struct {
	DB *sql.DB
}

// use pointer in one func = use it everywhere (can return nil)
func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	passwordHash := string(hashedPass)

	user := User{
		Email:         email,
		Password_hash: passwordHash,
	}

	row := us.DB.QueryRow(`INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}
