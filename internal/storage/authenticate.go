package storage

import (
	"database/sql"
	"fmt"

	"github.com/ruziba3vich/e_commerce_db/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username string, password string, db *sql.DB) (string, error) {
	var hashedPassword string
	var existingId int
	err := db.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&existingId, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	expirationHours := 24 * 30
	token, err := auth.GenerateJWT(existingId, username, expirationHours)
	if err != nil {
		return "", err
	}

	return token, nil
}
