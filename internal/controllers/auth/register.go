package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Register(ctx *gin.Context) {
	req := domain.RegisterUserReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}

	registerResp, err := c.userService.Register(ctx, req)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, registerResp, nil, true)
}
