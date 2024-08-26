package service

import (
	"github.com/lazbord/SpotyGo/common/model"
	"github.com/lazbord/SpotyGo/services/auth/database"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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
	}
	if !CheckPasswordHash(password, auth.UserPassword) {
		return "", errors.New("Wrong password")
	}

	return auth.ID, nil
}

func (a *AuthService) CreateUser() {
	auth := model.Auth{
		UserMail:     "lazarebordereaux@yahoo.fr",
		UserPassword: "password",
	}

	a.db.CreateAuth(auth)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
