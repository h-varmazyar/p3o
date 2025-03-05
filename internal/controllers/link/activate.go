package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller)Activate(ctx *gin.Context) {
	key:=ctx.Param("key")
	if err:=c.linkService.Activate(ctx, utils.FetchUserId(ctx), key); err!=nil{
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}
	utils.JsonHttpResponse(ctx, nil, nil, true)
}