package model

import "gorm.io/gorm"

//! after transaction request (paid), this cart will be deleted
type Cart struct {
	gorm.Model
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Total     int `json:"total"`
	Price     int `json:"price"`
}

type CartRequest struct {
	UserID int `json:"user_id"`
	CartID int `json:"cart_id"`
}

type CartData struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Carts  []Cart `json:"carts"`
}

/* NOTE
0. user will always receive their user_id from cookie/session/context value.
1. user will add cart by send array/json of cart struct and his user_id. (CartRequest struct)
2. user will remove cart by send array/json of cart_id and his user_id. (CartRequest struct)
3. user will update cart by send int/json of cart_id and his user_id. (CartRequest struct)
*/
