package user

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) SubmitVerificationCode(ctx *gin.Context) {
	req := domain.SubmitVerificationCodeReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}
	req.UserId = utils.FetchUserId(ctx)

	err := c.userService.SubmitVerificationCode(ctx, req)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, "", nil, true)
}
