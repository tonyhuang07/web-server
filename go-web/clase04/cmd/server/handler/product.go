package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/products"
)

type request struct {
	Name      string  `json:"name" `
	Price     float64 `json:"price" `
	Quality   int     `json:"quality"`
	Published bool    `json:"published"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p}
}

func (p *Product) Store(ctx *gin.Context) {
	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.Store(req.Name, req.Price, req.Quality, req.Published)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, product)
}

func (p *Product) GetAll(ctx *gin.Context) {
	products, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}
