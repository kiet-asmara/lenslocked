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

func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)

	user := User{
		Email: email,
	}

	row := us.DB.QueryRow(`
		SELECT id, password_hash
		FROM users 
		WHERE email = $1`, email)
	err := row.Scan(&user.ID, &user.Password_hash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	return &user, nil
}
