package user

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Edit(ctx *gin.Context) {
	editUserReq := domain.EditUserReq{}

	if err := ctx.ShouldBindJSON(&editUserReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}
	editUserReq.UserId = utils.FetchUserId(ctx)

	err := c.userService.EditUser(ctx, editUserReq)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, "", nil, true)
}
