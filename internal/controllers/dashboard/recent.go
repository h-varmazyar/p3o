package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/pagination"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Recent(ctx *gin.Context) {
	paginator := pagination.Paginator{
		Page:     1,
		PageSize: pagination.PageSize10,
	}
	links, err := c.linkSrv.List(ctx, utils.FetchUserId(ctx), paginator)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, links, nil, true)
}
