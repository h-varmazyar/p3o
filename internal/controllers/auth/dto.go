package auth

import "time"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token        string    `json:"token"`
	ExpireAt     time.Time `json:"expire_at"`
	VerifiedUser bool      `json:"verified_user"`
}
