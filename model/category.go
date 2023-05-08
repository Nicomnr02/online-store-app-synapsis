package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   int    `json:"id"`
	Type string `json:"category_type"`
}

// ! Category model for debugging
type CategoryRequest struct {
	Type string `json:"category_type"`
}

type CategoryData struct {
	ID       int       `json:"id"`
	Type     string    `json:"category_type"`
	Products []Product `json:"products"`
}
