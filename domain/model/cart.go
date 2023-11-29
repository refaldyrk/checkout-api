package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	UserID        string         `json:"user_id" gorm:"default:null"`
	IsCheckout    bool           `json:"is_checkout" gorm:"default:false"`
	CheckedOutAt  time.Time      `json:"checked_out_at" gorm:"default:null"`
	ProductOrders []ProductOrder `json:"product_orders"`
}
