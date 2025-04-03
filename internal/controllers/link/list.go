package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

type linksResp struct {
	Links      []domain.Link `json:"links"`
	Page       uint          `json:"page"`
	TotalPages uint          `json:"total_pages"`
}

func (c Controller) List(ctx *gin.Context) {
	links, err := c.linkService.List(ctx, utils.FetchUserId(ctx))
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	resp := linksResp{
		Links:      links.Links,
		Page:       1,
		TotalPages: 9,
	}
	utils.JsonHttpResponse(ctx, resp, nil, true)
}
