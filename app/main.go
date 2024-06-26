package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/ruziba3vich/e_commerce_db/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "e_commerce_db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbNames := []string{"users", "units", "categories", "products"}

	for _, dbName := range dbNames {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(string(sqlFile))
		fmt.Println(string(sqlFile))
		if err != nil {
			log.Fatal(err)
		}
	}

	var mtx *sync.Mutex

	router.GET("/getProduct/:productName", func(c *gin.Context) {
		productName := c.Param("productName")
		handlers.GetProductsByCategory(c, productName, db)
	})

	router.POST("/getProduct/:userId", func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		handlers.BuyProduct(c, db, userId, mtx)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
