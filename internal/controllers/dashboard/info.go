package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

type infoResp struct {
	TotalLink      domain.DashboardInfoItem `json:"total_link"`
	TotalVisit     domain.DashboardInfoItem `json:"total_visit"`
	DailyVisit     domain.DashboardInfoItem `json:"daily_visit"`
	UniqueVisitors domain.DashboardInfoItem `json:"unique_visitors"`
}

func (c Controller) Info(ctx *gin.Context) {
	resp := infoResp{}
	var err error
	if resp.TotalVisit, err = c.linkSrv.TotalVisit(ctx, utils.FetchUserId(ctx)); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}
	if resp.DailyVisit, err = c.linkSrv.TodayInfo(ctx, utils.FetchUserId(ctx)); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}
	if resp.TotalLink, err = c.linkSrv.TotalLinkCount(ctx, utils.FetchUserId(ctx)); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, resp, nil, true)
}
