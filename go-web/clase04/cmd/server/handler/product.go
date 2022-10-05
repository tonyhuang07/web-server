package handler

import (
	"strconv"

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

func (p *Product) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		ctx.JSON(400, gin.H{"error": "name is required"})
		return
	}

	if req.Price == 0 {
		ctx.JSON(400, gin.H{"error": "price is required"})
		return
	}

	if req.Quality == 0 {
		ctx.JSON(400, gin.H{"error": "quality is required"})
		return
	}

	if !req.Published {
		ctx.JSON(400, gin.H{"error": "published is required"})
		return
	}

	product, err := p.service.Update(id, req.Name, req.Price, req.Quality, req.Published)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, product)
}

func (p *Product) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := p.service.Delete(id); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, gin.H{"message": "product deleted"})
}

func (p *Product) UpdatePrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if req.Price == 0 {
		ctx.JSON(400, gin.H{"error": "price is required"})
		return
	}
	product, err := p.service.UpdatePrice(id, req.Price)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, product)
}
