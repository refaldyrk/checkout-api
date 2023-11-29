package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
	Carts     []Cart `json:"carts"`
}
