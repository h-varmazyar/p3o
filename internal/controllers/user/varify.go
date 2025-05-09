package user

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Verify(ctx *gin.Context) {
	verifyResp, err := c.userService.Verify(ctx, utils.FetchUserId(ctx))
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, verifyResp, nil, true)
}
