package model

import "time"

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	PhoneVerifiedAt time.Time `json:"phone_verified_at"`
	Password        string    `json:"password"`
	Gender          string    `json:"gender"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
