package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	BuyerID     uint    `json:"buyer_id"`
	TotalAmount float64 `json:"total_amount"`
	Currency    string  `json:"currency"`
	Status      string  `json:"status"`
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
