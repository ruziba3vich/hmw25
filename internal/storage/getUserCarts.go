package storage

import (
	"database/sql"
	"sync"

	"github.com/ruziba3vich/e_commerce_db/internal/models"
)

type CartDTO struct {
	Id       int                 `json:"id"`
	UserId   int                 `json:"user_id"`
	Products []models.ProductDTO `json:"products"`
}

// type ProductDTO struct {
// 	Id              int    `json:"id"`
// 	Name            string `json:"name"`
// 	Category        string `json:"category"`
// 	Price           int    `json:"price"`
// 	Unit            string `json:"unit"`
// 	Description     string `json:"description"`
// 	NumberOfProduct int    `json:"number_of_product"`
// }

func GetUserCarts(userId int, db *sql.DB, mtx *sync.Mutex) (results []models.ProductDTO, e error) {
	var user models.User
	getUserQuery := `
		SELECT
			id,
			name,
			balance,
			surname,
			username,
			password
		FROM Users
		WHERE id = $1;
	`
	err := db.QueryRow(getUserQuery, userId).Scan(&user.Id,
		&user.Name,
		&user.Balance,
		&user.Surname,
		&user.Username,
		&user.Password)
	query := `
		SELECT id, product_id FROM Carts WHERE user_id = $1;
	`
	rows, err := db.Query(query, userId)

	if err != nil {
		return nil, err
	}
	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, id := range ids {
		prds, err := user.BuyProducts(id, db, mtx)
		if err != nil {
			return nil, err
		}
		results = append(results, prds...)
	}
	return results, nil
}
