package user

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) GetInfo(ctx *gin.Context) {
	userId := utils.FetchUserId(ctx)

	userInfo, err := c.userService.Get(ctx, userId)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, userInfo, nil, true)
}
