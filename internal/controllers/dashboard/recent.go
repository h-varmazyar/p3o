package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Recent(ctx *gin.Context) {
	links, err := c.linkSrv.List(ctx, utils.FetchUserId(ctx))
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, links, nil, true)
}
