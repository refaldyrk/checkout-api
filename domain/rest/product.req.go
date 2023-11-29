package rest

import error_utils "checkout-api/utils/error"

type GetProductListReq struct {
	Category  string `form:"category" default:""`
	Search    string `form:"search" default:""` // using regex
	Page      int    `form:"page" default:"1"`
	Limit     int    `form:"limit" default:"10"`
	SortBy    string `form:"sort_by" default:"created_at"`
	SortOrder string `form:"sort_order" default:"desc"`
}

type PostAddItemToCartReq struct {
	Quantity int `json:"quantity" default:"1"`
}

func (req *PostAddItemToCartReq) Validate() error {
	if req.Quantity <= 0 {
		return &error_utils.CustomErr{
			Code:    400,
			Message: "quantity must be greater than 0",
		}
	}
	return nil
}
