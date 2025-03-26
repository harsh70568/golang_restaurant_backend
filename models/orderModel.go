package models

import "time"

type Order struct {
	ID         uint      `json:"order_id" gorm:"primaryKey;autoIncrement"`
	Order_date time.Time `json:"order_date" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	TableID    uint      `json:"table_id" validate:"required"`
	// Order_ID   string    `json:"order_id"`
}
