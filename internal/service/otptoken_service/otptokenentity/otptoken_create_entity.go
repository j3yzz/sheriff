package otptokenentity

import (
	"time"
)

type OtpTokenCreateEntity struct {
	ID       uint
	UserID   uint
	Token    string
	ExpireAt time.Time
}
