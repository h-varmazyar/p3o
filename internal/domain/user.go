package domain

import "time"

type ChangePasswordReq struct {
	UserId      uint   `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserInfo struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	IsVerified bool   `json:"is_verified"`
}

type EditUserReq struct {
	UserId    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type SubmitVerificationCodeReq struct {
	UserId uint   `json:"user_id"`
	Code   string `json:"code"`
	Mobile string `json:"mobile"`
}

type VerifyUserResp struct {
	ValidUntil time.Time `json:"valid_until"`
}
