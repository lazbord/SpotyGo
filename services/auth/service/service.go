package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func (a *AuthService) CheckCreditential(email, password string) (string, *model.Auth, error) {
	auth, err := a.db.GetAuthByEmail(email)
	if err != nil {
		return "", nil, errors.New("No user with this email")
	}
	if !CheckPasswordHash(password, auth.UserPassword) {
		return "", nil, errors.New("Wrong password")
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	token, err := generateToken(auth.ID, &expirationTime)
	if err != nil {
		return "", nil, err
	}
	return token, auth, nil
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

func generateToken(userid string, exp *time.Time) (string, error) {
	key := os.Getenv("JWT_KEY")
	var claims *model.Claims
	if exp != nil {
		claims = &model.Claims{
			UserID: userid,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(*exp),
			},
		}
	} else {
		claims = &model.Claims{
			UserID: userid,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signed, nil
}
