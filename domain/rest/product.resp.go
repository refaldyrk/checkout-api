package rest

import (
	"checkout-api/domain/model"
	"time"
)

type GetProductListResp struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func (slf *GetProductListResp) ParseFromEntityList(products []model.Product) []GetProductListResp {
	var result []GetProductListResp
	for _, product := range products {
		parsed := GetProductListResp{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			Name:      product.Name,
			Price:     product.Price,
			Category:  product.Category,
		}
		if !product.DeletedAt.Valid {
			parsed.DeletedAt = time.Time{}
		} else {
			parsed.DeletedAt = product.DeletedAt.Time
		}
		result = append(result, parsed)
	}

	return result
}
