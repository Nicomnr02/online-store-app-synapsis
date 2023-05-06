package model

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	SessionToken string `json:"session_token"`
}
