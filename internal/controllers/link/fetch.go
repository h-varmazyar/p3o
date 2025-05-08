package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Fetch(ctx *gin.Context) {
	key := ctx.Param("key")

	if details, err := c.linkService.Fetch(ctx, utils.FetchUserId(ctx), key); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
	} else {
		utils.JsonHttpResponse(ctx, details, nil, true)
	}
}
