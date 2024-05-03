package handlers

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/e_commerce_db/internal/models"
)

type CartId struct {
	CartId int `json:"cart_id"`
}

func BuyProduct(c *gin.Context, db *sql.DB, userId int, mtx *sync.Mutex) {
	var cartId CartId
	err := c.ShouldBindJSON(&cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	query := `
		SELECT
			name,
			balance,
			surname,
			username,
			password
		FROM Users
		WHERE id = $1;
	`
	var user models.User
	err = db.QueryRow(query, userId).Scan(&user.Name, &user.Balance, &user.Surname, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user.Id = userId
	products, err := user.BuyProducts(cartId.CartId, db, mtx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, products)
}

/*

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Balance  uint   `json:"balance"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}


*/
