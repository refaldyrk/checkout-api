package handler

import (
	"checkout-api/domain/rest"
	"checkout-api/service"
	"checkout-api/utils/http_response"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	respWriter  http_response.IResponseWriter
	cartService service.ICartService
}

type ICartHandler interface {
	GetCartItems(ctx *gin.Context)
	RemoveItemFromCart(ctx *gin.Context)
	CheckoutCart(ctx *gin.Context)
}

func NewCartHandler(respWriter http_response.IResponseWriter, cartService service.ICartService) ICartHandler {
	return &CartHandler{
		respWriter:  respWriter,
		cartService: cartService,
	}
}

// @Summary get product list from current cart (MANDATORY)
// @Tags Cart
// @Router /cart/orders [get]
// @Security BearerAuth
// @Success 200 {object} rest.BaseJSONResp{}
func (slf *CartHandler) GetCartItems(ctx *gin.Context) {
	user_id := ctx.GetString("user_id")
	orders, err := slf.cartService.GetCartItems(user_id)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	dto := rest.GetCartItemsResp{}
	resp := []rest.GetCartItemsResp{}
	resp = append(resp, dto.ParseFromEntityList(orders)...)

	slf.respWriter.HTTPJson(ctx, resp, "OK")
}

// @Summary remove item from cart (MANDATORY)
// @Tags Cart
// @Router /cart/orders/{id} [delete]
// @Security BearerAuth
// @Param id path string true "order id"
// @Success 200 {object} rest.BaseJSONResp{}
func (slf *CartHandler) RemoveItemFromCart(ctx *gin.Context) {
	order_id := ctx.Param("id")
	user_id := ctx.GetString("user_id")

	err := slf.cartService.RemoveItemFromCart(order_id, user_id)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, nil, "item removed")
}

// @Summary checkout cart (MANDATORY)
// @Tags Cart
// @Router /cart/checkout [post]
// @Security BearerAuth
// @Param payload body rest.PostCheckoutCartReq true "payload"
// @Success 200 {object} rest.BaseJSONResp{data=rest.PostCheckoutCartResp}
func (slf *CartHandler) CheckoutCart(ctx *gin.Context) {
	user_id := ctx.GetString("user_id")

	var payload rest.PostCheckoutCartReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	moneyReturn, err := slf.cartService.CheckoutCart(user_id, payload.MoneyInput)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, rest.PostCheckoutCartResp{
		MoneyReturn: moneyReturn,
	}, "OK")
}
