package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Create(ctx *gin.Context) {
	linkCreateReq := new(domain.LinkCreateReq)

	if err := ctx.ShouldBindJSON(&linkCreateReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}

	linkCreateReq.UserId=utils.FetchUserId(ctx)

	resp, err:=c.linkService.Create(ctx, *linkCreateReq)
	if err!=nil{
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	utils.JsonHttpResponse(ctx, resp, nil, true)
}