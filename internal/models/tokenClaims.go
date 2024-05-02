package models

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	jwt.StandardClaims
}
