package storage

import (
	"database/sql"

	"github.com/ruziba3vich/e_commerce_db/internal/models"
)

func GetProductsByCategory(categoryName string, db *sql.DB) ([]models.ProductDTO, error) {
	query := `
	SELECT EXISTS (
		SELECT 1
		FROM categories
		WHERE name = $1
	);
	`
	row := db.QueryRow(query, categoryName)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		// log.Println("came --------------")
		return nil, err
	}

	if exists {
		query = `
			SELECT
				p.id,
				p.name,
				c.name AS category_name,
				p.price,
				u.name AS unit_name,
				p.description,
				p.number_of_product
			FROM Products p 
			INNER JOIN Categories c ON c.id = p.category_id
			INNER JOIN Units u ON u.id = p.unit_id
			WHERE c.name = $1;
		`

		rows, err := db.Query(query, categoryName)
		if err != nil {
			// log.Println("came --------------")
			return nil, err
		}
		defer rows.Close()

		var products []models.ProductDTO
		for rows.Next() {
			var product models.ProductDTO
			err := rows.Scan(&product.Id,
				&product.Name,
				&product.Category,
				&product.Price,
				&product.Unit,
				&product.Description,
				&product.NumberOfProduct)
			if err != nil {
				// log.Println("came --------------")
				return nil, err
			}
			products = append(products, product)
		}
		return products, nil
	}
	return []models.ProductDTO{}, nil
}
