package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) All(ctx *gin.Context) {
	fmt.Println("user is:", utils.FetchUserId(ctx))
	links, err := c.linkService.All(ctx, utils.FetchUserId(ctx))
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, links, nil, true)
}
