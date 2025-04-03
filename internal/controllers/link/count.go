package link

import (
	"github.com/gin-gonic/gin"
)

type countsResp struct {
	TotalLinks  uint   `json:"total_links"`
	Growth      string `json:"growth"`
	GrowthTrend string `json:"growth_trend"`
}

func (c Controller) Counts(ctx *gin.Context) {
	//if totalLinkCount, err := c.linkService.TotalLinkCount(ctx, utils.FetchUserId(ctx)); err != nil {
	//	utils.JsonHttpResponse(ctx, nil, err, false)
	//} else {
	//	resp := countsResp{
	//		Growth:      fmt.Sprintf("%.1f%", 3.22341),
	//		GrowthTrend: "+",
	//		TotalLinks:  totalLinkCount,
	//	}
	//	utils.JsonHttpResponse(ctx, resp, nil, true)
	//}
}
