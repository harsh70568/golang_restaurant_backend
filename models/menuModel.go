package models

import "time"

type Menu struct {
	ID         uint      `json:"menu_id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	// MenuID     string    `json:"menu_id"`
}
