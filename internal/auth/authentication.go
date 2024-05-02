package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ruziba3vich/e_commerce_db/internal/models"
)

var secretKey = []byte("f350124f5d1497b709735c3cc842041b3534c0e1a87f0ba0abd5c37ee8369bf4")

func GenerateJWT(userID int, username string, expirationHours int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	claims := &models.TokenClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func IsTokenExpired(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return true
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	return time.Now().After(expirationTime)
}

func IsTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	return !time.Now().After(expirationTime)
}

func GetJwtToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

