package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	dbNames := []string{"units", "categories", "products"}

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

	router.POST("/getProduct/:productName", func(c *gin.Context) {
		productName := c.Param("productName")
		handlers.GetProductsByCategory(c, productName, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
