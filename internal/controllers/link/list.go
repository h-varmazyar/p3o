package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/pagination"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) List(ctx *gin.Context) {
	links, err := c.linkService.List(ctx, utils.FetchUserId(ctx), pagination.GinPaginator(ctx))
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, links, nil, true)
}
