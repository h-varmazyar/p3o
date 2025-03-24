package domain

import "time"

type Tokens struct {
	Token      string    `json:"token"`
	ExpireAt   time.Time `json:"expire_at"`
	IsVerified bool      `json:"is_verified"`
}
