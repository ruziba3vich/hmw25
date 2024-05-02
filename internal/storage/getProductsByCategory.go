package storage

import (
	"database/sql"
)

func GetProductsByCategory(categoryName string, db *sql.DB) ([]ProductDTO, error) {
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
			SELECT p.id, p.name, c.name AS category_name, p.price, u.name AS unit_name, p.description
			FROM Products p 
			INNER JOIN Categories c ON c.id = p.category_id
			INNER JOIN Units u ON u.id = p.unit_id;
		`

		rows, err := db.Query(query)
		if err != nil {
			// log.Println("came --------------")
			return nil, err
		}
		defer rows.Close()

		var products []ProductDTO
		for rows.Next() {
			var product ProductDTO
			err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Unit, &product.Description)
			if err != nil {
				// log.Println("came --------------")
				return nil, err
			}
			products = append(products, product)
		}
		return products, nil
	}
	return []ProductDTO{}, nil
}

type ProductDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}
