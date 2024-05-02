package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ruziba3vich/e_commerce_db/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "e_commerce_db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbNames := []string{}

	for _, dbName := range dbNames {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			log.Fatal(err)
		}
	}

	router.POST("/twit", func(c *gin.Context) {
		handlers.CreateTwit(c, db)
	})

	router.POST("/comment/:id", func(c *gin.Context) {
		handlers.Comment(c, db, context.Background())
	})

	router.POST("/", func(c *gin.Context) {
		handlers.LoadTwits(c, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
