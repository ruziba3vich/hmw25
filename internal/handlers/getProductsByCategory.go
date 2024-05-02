package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/e_commerce_db/internal/storage"
)

func GetProductsByCategory(c *gin.Context, categoryName string, db *sql.DB) {
	products, err := storage.GetProductsByCategory(categoryName, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("the category %s does not exists", categoryName)})
		return
	}
	c.JSON(http.StatusOK, products)
}
