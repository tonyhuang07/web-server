package main

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const ValidToken = "123456"

type Product struct {
	ID        int     `json:"id" `
	Name      string  `json:"name" `
	Price     float64 `json:"price" `
	Quality   int     `json:"quality"`
	Published bool    `json:"published" binding:"required"`
}

func Create(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != ValidToken {
		ctx.JSON(401, gin.H{"error": "invalid token, no permission to make this request"})
		return
	}

	var product Product
	var emptyFiled []string
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	product.ID = len(products) + 1
	values := reflect.ValueOf(product)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == reflect.Zero(values.Field(i).Type()).Interface() {
			emptyFiled = append(emptyFiled, values.Type().Field(i).Name)
		}
	}
	if len(emptyFiled) > 0 {
		ctx.JSON(400, gin.H{"error": "The following fields: " + strings.Join(emptyFiled, " & ") + " are required"})
		return
	}

	products = append(products, product)
	ctx.JSON(200, product)
}

var products []Product

func main() {
	r := gin.Default()
	r.POST("/products", Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
