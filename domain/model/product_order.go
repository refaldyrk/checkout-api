package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductOrder struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	CartID    string  `json:"cart_id" gorm:"default:null"`
	ProductID string  `json:"product_id" gorm:"default:null"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
}
