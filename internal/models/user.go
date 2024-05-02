package models

import (
	"database/sql"
	"errors"
	"log"
	"sync"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Balance  uint   `json:"balance"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Balance  uint   `json:"balance"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
}

func (u *User) AddProductIntoCart(productId, numberOfProduct int, db *sql.DB) (e error) {
	query := `
		INSERT INTO Cart(
			user_id,
			product_id,
			number_of_product
		) VALUES ($1, $2, $3)
		RETURNING id;
	`

	_, err := db.Exec(query, u.Id, numberOfProduct)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) RemoveProductFromCart(productId int, db *sql.DB) (bool, error) {
	query := `
		DELETE FROM Cart c
		WHERE c.user_id = $1 AND c.product_id = $2
	`

	result, err := db.Exec(query, u.Id, productId)

	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, errors.New("no rows has been deleted")
	}
	return true, nil
}

/// <------------------------------ the part I used transaction --------------------------------------->

func (u *User) BuyProducts(cartId int, db *sql.DB, mtx *sync.Mutex) (products []ProductDTO, e error) {
	var totalProductsPrice int

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("failed to begin transaction,", err)
	}
	defer tx.Rollback()

	query := `
		SELECT
			p.id,
			p.name,
			c.name AS category,
			p.price,
			u.name AS unit,
			p.description,
			p.number_of_product
		FROM Cart cart
		INNER JOIN Products p ON cart.product_id = p.id
		INNER JOIN Categories c ON c.id = p.category_id
		INNER JOIN Units u ON u.id = p.unit_id;
	`

	rows, err := tx.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product ProductDTO
		rows.Scan(&product.Id,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.Unit,
			&product.Description,
			&product.NumberOfProduct)
		products = append(products, product)
		totalProductsPrice += product.Price
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if totalProductsPrice > int(u.Balance) {
		return nil, errors.New("check your balance")
	}

	for _, product := range products {
		u.Balance -= uint(product.Price)
	}
	query = `
		UPDATE Users u
		SET u.balance = $1
		WHERE u.user_id = $2;
	`
	_, err = tx.Exec(query, u.Balance, u.Id)
	if err != nil {
		return nil, err
	}

	query = `
		UPDATE Products p
		SET p.number_of_product = p.number_of_product - $1
		WHERE p.id = $2;
	`
	var wg sync.WaitGroup
	for _, product := range products {
		wg.Add(1)
		go func(product ProductDTO) {
			defer wg.Done()
			mtx.Lock()
			_, err := tx.Exec(query, product.NumberOfProduct, product.Id)
			mtx.Unlock()
			if err != nil {
				log.Fatal(err)
				return
			}
		}(product)
	}
	wg.Wait()
	tx.Commit()
	return products, nil
}
