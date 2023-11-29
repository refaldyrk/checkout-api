package rest

type PostCheckoutCartReq struct {
	MoneyInput float64 `json:"money_input" binding:"required"`
}
