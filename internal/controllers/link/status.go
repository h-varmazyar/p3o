package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller)Status(ctx *gin.Context){
	key:=ctx.Param("key")

	if status, err:=c.linkService.Status(ctx, utils.FetchUserId(ctx), key); err!=nil{
		utils.JsonHttpResponse(ctx, nil, err, false)
	}else{
		utils.JsonHttpResponse(ctx, map[string]interface{}{"status": status}, nil, true)
	}
}