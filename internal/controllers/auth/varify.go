package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Verify(ctx *gin.Context) {
	req := domain.VerifyUserReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}
	token, err := c.userService.Verify(ctx, req)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, token, nil, true)
}
