package link

import (
	"github.com/gin-gonic/gin"
)

type visitsResp struct {
	Visits      uint   `json:"visits"`
	Growth      string `json:"growth"`
	GrowthTrend string `json:"growth_trend"`
}

func (c Controller) Visits(ctx *gin.Context) {
	//if totalVisits, err := c.linkService.TotalVisits(ctx, utils.FetchUserId(ctx)); err != nil {
	//	utils.JsonHttpResponse(ctx, nil, err, false)
	//} else {
	//	resp := visitsResp{
	//		Visits:      totalVisits,
	//		Growth:      fmt.Sprintf("%.1f%", 43.441),
	//		GrowthTrend: "+",
	//	}
	//	utils.JsonHttpResponse(ctx, resp, nil, true)
	//}
}
