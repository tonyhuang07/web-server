package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

const FilePath = "products.json"

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Code        string `json:"code"`
	Published   bool   `json:"published"`
	CreatedDate string `json:"created_date"`
}

func getAll() gin.HandlerFunc {
	content, err := os.ReadFile(FilePath)
	if err != nil {
		return func(ctx *gin.Context) {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}
	var payload []Product
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return func(ctx *gin.Context) {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}
	return func(ctx *gin.Context) {
		ctx.JSON(200, payload)
	}

}

func main() {
	r := gin.Default()
	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Tony!",
		})
	})

	r.GET("/products", getAll())
	r.Run()
}
