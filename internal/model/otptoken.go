package model

import "time"

type OtpToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `json:"token"`
	UserID    uint      `json:"user_id"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"references:ID"`
}
