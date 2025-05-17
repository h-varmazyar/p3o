package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c Controller) Edit(ctx *gin.Context) {
	editLinkReq := domain.EditLinkReq{}

	if err := ctx.ShouldBindJSON(&editLinkReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}

	editLinkReq.Key = ctx.Param("key")
	editLinkReq.UserId = utils.FetchUserId(ctx)

	fmt.Println(editLinkReq)
	if err := c.linkService.Edit(ctx, editLinkReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
	} else {
		utils.JsonHttpResponse(ctx, "", nil, true)
	}
}
