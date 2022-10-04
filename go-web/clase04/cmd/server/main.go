package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyhuang07/web-server/go-web/clase04/cmd/server/handler"
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/products"
)

func main() {
	r := gin.Default()
	repository := products.NewRepository()
	service := products.NewService(repository)
	product := handler.NewProduct(service)
	r.POST("/products", product.Store)
	r.GET("/products", product.GetAll)
	r.Run()
}
