package storage

import (
	"database/sql"
	"fmt"

	"github.com/ruziba3vich/e_commerce_db/internal/auth"
	"github.com/ruziba3vich/e_commerce_db/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User, db *sql.DB) (string, error) {
	var existingUsername string
	err := db.QueryRow("SELECT username FROM users WHERE username = $1", user.Username).Scan(&existingUsername)
	if err != sql.ErrNoRows {
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	query := `
	INSERT INTO users (
		username,
		password)
	VALUES ($1, $2)
	RETURNING id;
	`
	err = db.QueryRow(query, user.Username, string(hashedPassword)).Scan(&user.Id)
	if err != nil {
		return "", err
	}

	expirationHours := 24 * 30
	token, err := auth.GenerateJWT(user.Id, existingUsername, expirationHours)
	if err != nil {
		return "", err
	}

	return token, nil
}
