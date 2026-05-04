package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hernanprovoste/icetrack-pro/internal/core/product"
	"github.com/hernanprovoste/icetrack-pro/internal/db"
)

func main() {

	cfg := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "icetrack_pro",
		SSLMode:  "disable",
	}

	db, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("failed to auto migrate database: %v", err)
	}

	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		products, err := product.ListAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list products"})
			return
		}

		summaries := make([]product.ProductSummary, 0, len(products))
		for _, p := range products {
			summaries = append(summaries, product.ProductSummary{
				ID:     p.ID,
				Name:   p.Name,
				Price:  p.Price,
				Active: p.Active,
			})
		}

		c.JSON(http.StatusOK, summaries)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
			return
		}

		p, err := product.FindByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not find product"})
			return
		}

		if p == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		c.JSON(http.StatusOK, p)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
