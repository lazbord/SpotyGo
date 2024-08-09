package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	ID           string `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       string `bson:"user_id" json:"user_id"`
	UserMail     string `bson:"user_email" json:"user_mail"`
	UserPassword string `bson:"user_password" json:"user_password"`
}

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}
