package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Delete(ctx *gin.Context) {
	key:=ctx.Param("key")

	if err:=c.linkService.Delete(ctx, utils.FetchUserId(ctx), key); err!=nil{
		utils.JsonHttpResponse(ctx, nil, err, false)
	}else{
		utils.JsonHttpResponse(ctx, nil, nil, true)
	}
}