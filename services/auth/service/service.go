package service

import (
	"github.com/lazbord/SpotyGo/services/auth/database"
	"github.com/lazbord/SpotyGo/services/auth/model"
	"github.com/pkg/errors"
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
	auth, err := a.db.GetAuthByEmail(email)
	if err != nil {
		return "", errors.New("No user with this email")
	} else if auth.UserPassword != password {
		return "", errors.New("Wrong password")
	}

	return auth.UserID, nil
}

func (a *AuthService) CreateUser() {
	auth := model.Auth{
		UserMail:     "lazarebordereaux@yahoo.fr",
		UserPassword: "password",
	}
	a.db.CreateAuth(auth)
}
