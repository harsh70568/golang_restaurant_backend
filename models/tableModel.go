package models

import "time"

type Table struct {
	ID           uint      `json:"table_id" gorm:"primaryKey;autoIncrement"`
	Guests       int       `json:"guests" validate:"required"`
	Table_number int       `json:"table_number" validate:"required"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	// TableID      string    `json:"table_id"`
}
