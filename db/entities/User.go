package entities

import "github.com/dgrijalva/jwt-go"

type User struct {
	Name     string    `json:"name" bson:"name,omitempty"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
	SID      jwt.Token `json:"sid" bson:"sid"`
}
