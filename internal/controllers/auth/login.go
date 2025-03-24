package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"time"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token        string    `json:"token"`
	ExpireAt     time.Time `json:"expire_at"`
	VerifiedUser bool      `json:"verified_user"`
}

func (c *Controller) Login(ctx *gin.Context) {
	loginReq := new(LoginReq)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrLoginFailed.AddOriginalError(err), false)
		return
	}

	tokens, err := c.userService.Login(ctx, loginReq.Username, loginReq.Password)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	resp := LoginResp{
		Token:        tokens.Token,
		ExpireAt:     tokens.ExpireAt,
		VerifiedUser: tokens.IsVerified,
	}

	utils.JsonHttpResponse(ctx, resp, nil, true)
}
