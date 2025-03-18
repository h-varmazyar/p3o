package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Visits(ctx *gin.Context) {
	if totalVisits, err := c.linkService.TotalVisits(ctx, utils.FetchUserId(ctx)); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
	} else {
		utils.JsonHttpResponse(ctx, map[string]interface{}{"visits": totalVisits}, nil, true)
	}
}
