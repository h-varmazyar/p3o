package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) IndirectVisit(ctx *gin.Context) {
	key := ctx.Param("key")
	id := ctx.Param("id")

	link, err := c.linkService.IndirectVisit(ctx, key, id)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, link, nil, true)

}
