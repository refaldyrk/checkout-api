package rest

import (
	"checkout-api/domain/model"
	"time"
)

type GetCartItemsResp struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	CartID    string        `json:"cart_id" gorm:"default:null"`
	ProductID string        `json:"product_id" gorm:"default:null"`
	Product   model.Product `json:"product"`
	Quantity  int           `json:"quantity"`
}

func (slf *GetCartItemsResp) ParseFromEntityList(orders []model.ProductOrder) []GetCartItemsResp {
	var result []GetCartItemsResp
	for _, order := range orders {
		parsed := GetCartItemsResp{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			CartID:    order.CartID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Product:   order.Product,
		}
		if !order.DeletedAt.Valid {
			parsed.DeletedAt = time.Time{}
		} else {
			parsed.DeletedAt = order.DeletedAt.Time
		}
		result = append(result, parsed)
	}

	return result
}

type PostCheckoutCartResp struct {
	MoneyReturn float64 `json:"money_return"`
}
