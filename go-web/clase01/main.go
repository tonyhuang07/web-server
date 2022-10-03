package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

const FilePath = "products.json"

// type Interval struct {
// 	GreaterThan int `json:"greater_than"`
// 	LessThan    int `json:"less_than"`
// }

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

func GetAll(ctx *gin.Context) {
	file, err := os.Open(FilePath)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	price, havePrice := ctx.GetQuery("price")
	if havePrice {
		priceInterval := map[string]int{}
		err := json.Unmarshal([]byte(price), &priceInterval)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var filteredProducts []Product
		for _, product := range products {
			if product.Price >= priceInterval["gt"] && product.Price <= priceInterval["lt"] {
				filteredProducts = append(filteredProducts, product)
			}
		}
		ctx.JSON(200, filteredProducts)
		return
	}
	ctx.JSON(200, []Product{})
	// published, haveFilter := ctx.GetQuery("published")
	// var filteredProducts []Product
	// if haveFilter {
	// 	for _, product := range products {
	// 		parsedPublished, err := strconv.ParseBool(published)
	// 		if err != nil {
	// 			ctx.JSON(500, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		if product.Published == parsedPublished {
	// 			filteredProducts = append(filteredProducts, product)
	// 		}
	// 	}
	// }

	// ctx.JSON(200, filteredProducts)
}
func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	file, err := os.Open(FilePath)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, product := range products {
		if product.ID == id {
			ctx.JSON(200, product)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "product not found"})

}

//	func GetTest(ctx *gin.Context) {
//		id, haveID := ctx.GetQuery("id")
//		if haveID {
//			log.Println("id", id)
//		}
//		name, haveName := ctx.GetQuery("name")
//		if haveName {
//			log.Println("name", name)
//		}
//	}
func main() {
	r := gin.Default()
	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Tony!",
		})
	})

	r.GET("/products", GetAll)
	r.GET("/products/:id", GetById)
	// r.GET("/test", GetTest)
	r.Run()
}
