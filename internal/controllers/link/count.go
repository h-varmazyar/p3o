package link

import(
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Counts(ctx *gin.Context) {
	if totalLinkCount, err:=c.linkService.TotalLinkCount(ctx, utils.FetchUserId(ctx)); err!=nil{
		utils.JsonHttpResponse(ctx, nil, err, false)
	} else {
		utils.JsonHttpResponse(ctx, totalLinkCount, nil, true)
	}
}