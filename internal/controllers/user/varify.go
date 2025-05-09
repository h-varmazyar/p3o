package user

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
	req.UserId = utils.FetchUserId(ctx)
	verifyResp, err := c.userService.Verify(ctx, req)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, verifyResp, nil, true)
}
