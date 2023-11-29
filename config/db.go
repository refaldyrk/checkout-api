package config

import (
	"checkout-api/domain/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(Envs.DB_URI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.ProductOrder{})
	return db, nil
}
