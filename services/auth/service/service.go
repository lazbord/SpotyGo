package service

import (
	"errors"

	"github.com/lazbord/SpotyGo/services/auth/database"
)

type AuthService struct {
	db *database.Adapter
}

func NewAuthService(db *database.Adapter) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CheckCreditential(email, password string) (string, error) {
	if email != "lazareb" || password != "pwd" {
		return "", errors.New("invalid email or password")
	}
	return "0123456789", nil
}

func (a *AuthService) CreateUser() {
	a.db.NewUser()
}
