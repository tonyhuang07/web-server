package handler

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/products"
	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/web"
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
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	if req.Name == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "name is required"))
		return
	}

	if req.Price == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "price is required"))
		return
	}

	if req.Quality == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "quality is required"))
		return
	}

	if !req.Published {
		ctx.JSON(400, web.NewResponse(400, nil, "published is required"))
		return
	}
	product, err := p.service.Store(req.Name, req.Price, req.Quality, req.Published)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, product)
}

func (p *Product) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
		return
	}
	products, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, products, ""))
}

func (p *Product) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "invalid id"))
		return
	}

	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "name is required"))
		return
	}

	if req.Price == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "price is required"))
		return
	}

	if req.Quality == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "quality is required"))
		return
	}

	if !req.Published {
		ctx.JSON(400, web.NewResponse(400, nil, "published is required"))
		return
	}

	product, err := p.service.Update(id, req.Name, req.Price, req.Quality, req.Published)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, product)
}

func (p *Product) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "invalid id"))
		return
	}

	if err := p.service.Delete(id); err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(204, web.NewResponse(204, nil, ""))
}

func (p *Product) UpdatePrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "invalid id"))
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	if req.Price == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "price is required"))
		return
	}
	product, err := p.service.UpdatePrice(id, req.Price)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, product, ""))
}
