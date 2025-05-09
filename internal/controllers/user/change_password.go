package user

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) ChangePassword(ctx *gin.Context) {
	changePasswordReq := domain.ChangePasswordReq{}

	if err := ctx.ShouldBindJSON(&changePasswordReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}
	changePasswordReq.UserId = utils.FetchUserId(ctx)

	err := c.userService.ChangePassword(ctx, changePasswordReq)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, "", nil, true)
}
