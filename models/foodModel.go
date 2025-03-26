package models

import "time"

type Food struct {
	ID         uint      `json:"food_id" gorm:"primaryKey;autoincrement"`
	Name       string    `json:"name" validate:"required,min=2,max=20"`
	Price      float32   `json:"price" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	MenuID     uint      `json:"menu_id" validate:"required"`
	// FoodID     string    `json:"food_id"`
}
