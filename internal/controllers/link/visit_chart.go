package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) VisitChart(ctx *gin.Context) {
	key := ctx.Param("key")
	resp, err := c.linkService.VisitChart(ctx, utils.FetchUserId(ctx), key)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, map[string]interface{}{"visits": resp}, nil, true)
}
