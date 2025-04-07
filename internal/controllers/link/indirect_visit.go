package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"net/http"
)

func (c Controller) IndirectVisit(ctx *gin.Context) {
	key := ctx.Param("key")
	id := ctx.Param("id")

	fmt.Println("sssssssssssssssssssssss")
	link, err := c.linkService.IndirectVisit(ctx, key, id)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, link.Url)
}
