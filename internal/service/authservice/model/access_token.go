package model

import "time"

type AccessToken struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
