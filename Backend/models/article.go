package models

import "gorm.io/gorm"

// ExchangeRate represents an exchange rate between two currencies
type Article struct {
	gorm.Model
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Preview string `json:"preview"`
	Likes   int    `json:"likes"`
}
