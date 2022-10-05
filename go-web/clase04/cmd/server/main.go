package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tonyhuang07/web-server/go-web/clase04/cmd/server/handler"
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/products"
	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	database := store.New(store.FileType, "products.json")
	repository := products.NewRepository(database)
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	rootProduct := r.Group("/products")
	rootProduct.POST("/", product.Store)
	rootProduct.GET("/", product.GetAll)
	rootProduct.PUT("/:id", product.Update)
	rootProduct.PATCH("/:id", product.UpdatePrice)
	rootProduct.DELETE("/:id", product.Delete)
	r.Run()
}
