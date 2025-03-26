package models

import "time"

type OrderItem struct {
	ID         uint      `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	Quantity   int       `json:"quantity" validate:"required"`
	UnitPrice  int       `json:"unit_price" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	FoodID     uint      `json:"food_id" validate:"required"`
	OrderID    uint      `json:"order_id" validate:"required"`
	// OrderItemID string    `json:"order_item_id"`
}
