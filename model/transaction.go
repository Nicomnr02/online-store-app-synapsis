package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	CartID int    `json:"cart_id"`
	Status string `json:"transaction_status"` // paid or not paid
}

type TransactionRequest struct {
	ID     int `json:"transaction_id"`
	UserID int `json:"user_id"`
	CartID int `json:"cart_id"`
}

type TransactionsRequest struct {
	UserID int `json:"user_id"`
	Carts []Cart `json:"all_carts"`
}
