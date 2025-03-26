package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName     string    `json:"first_name" validate:"required,min=2,max=20"`
	LastName      string    `json:"last_name" validate:"required,min=2,max=20"`
	Password      string    `json:"password" validate:"required,min=8,max=16"`
	Email         string    `json:"email" validate:"email,required" gorm:"unique"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Created_at    time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
