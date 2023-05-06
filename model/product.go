package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"product_name"`
	Price      int    `json:"product_price"`
	Stock      int    `json:"product_stock"`
}
