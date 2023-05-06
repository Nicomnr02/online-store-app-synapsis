package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID     int    `json:"id"`
	CartID int    `json:"cart_id"`
	Status string `json:"transaction_status"` // paid or not paid
}

type TransactionRequest struct {
	ID     int `json:"id"`
	CartID int `json:"cart_id"` 
	Cash   int `json:"transaction_cash"`
}

