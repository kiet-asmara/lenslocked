package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token only set when creating new session
	// when looking up session, will be empty
	// database only stores hash of the token, it cant be reversed
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(UserID int) (*Session, error) {
	// TODO:create session token
	// TODO: Implement sessionservice.create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement sessionservice.user
	return nil, nil
}
