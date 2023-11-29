package handler

import (
	"checkout-api/domain/rest"
	"checkout-api/service"
	"checkout-api/utils/http_response"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	respWriter     http_response.IResponseWriter
	productService service.IProductService
}

type IProductHandler interface {
	GetList(ctx *gin.Context)
	GetCategoryList(ctx *gin.Context)
	AddItemToCart(ctx *gin.Context)
}

func NewProductHandler(respWriter http_response.IResponseWriter, productService service.IProductService) IProductHandler {
	return &ProductHandler{
		respWriter:     respWriter,
		productService: productService,
	}
}

// @Summary get product list (MANDATORY)
// @Tags Product
// @Success 200 {object} rest.BaseJSONResp{data=[]rest.GetProductListResp}
// @Router /products [get]
// @param query  query  rest.GetProductListReq  false "query"
// @Security BearerAuth
func (slf *ProductHandler) GetList(ctx *gin.Context) {
	var query rest.GetProductListReq
	err := ctx.Bind(&query)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	products, err := slf.productService.GetList(query.Category, query.Search, query.Page, query.Limit, query.SortBy, query.SortOrder)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}
	result := make([]rest.GetProductListResp, 0)
	dto := rest.GetProductListResp{}
	result = append(result, dto.ParseFromEntityList(products)...)

	slf.respWriter.HTTPJson(ctx, result, "OK")
}

// @Summary get product category list
// @Tags Product
// @Success 200 {object} rest.BaseJSONResp{data=[]string}
// @Router /products/category-list [get]
// @Security BearerAuth
func (slf *ProductHandler) GetCategoryList(ctx *gin.Context) {
	categories, err := slf.productService.GetCategoryList()
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, categories, "OK")
}

// @Summary add item to cart (MANDATORY)
// @Tags Product
// @Router /products/{id}/add-to-cart [post]
// @Security BearerAuth
// @param id path string true "id"
// @param payload  body  rest.PostAddItemToCartReq  true "payload"
// @Success 200 {object} rest.BaseJSONResp{}
func (slf *ProductHandler) AddItemToCart(ctx *gin.Context) {
	product_id := ctx.Param("id")
	user_id := ctx.GetString("user_id")

	var payload rest.PostAddItemToCartReq
	err := ctx.BindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	err = payload.Validate()
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	err = slf.productService.AddItemToCart(product_id, payload.Quantity, user_id)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, nil, "item added")
}
