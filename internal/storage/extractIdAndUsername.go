package storage

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/ruziba3vich/e_commerce_db/internal/auth"
)

func ExtractUserIDAndUsername(tokenString string) (int, string, error) {
	if !auth.IsTokenValid(tokenString) {
		return 0, "", fmt.Errorf("invalid or expired token")
	}

	token, err := auth.GetJwtToken(tokenString)
	if err != nil {
		return 0, "", fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	userID := int(claims["userID"].(float64))
	username := claims["username"].(string)

	return userID, username, nil
}
