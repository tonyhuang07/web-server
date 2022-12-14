package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tonyhuang07/web-server/go-web/clase04/cmd/server/handler"
	"github.com/tonyhuang07/web-server/go-web/clase04/docs"
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/products"
	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/store"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootProduct := r.Group("/products")
	rootProduct.POST("/", product.Store)
	rootProduct.GET("/", product.GetAll)
	rootProduct.PUT("/:id", product.Update)
	rootProduct.PATCH("/:id", product.UpdateName)
	rootProduct.DELETE("/:id", product.Delete)
	errRun := r.Run()
	if errRun != nil {
		panic(err)
	}
}
