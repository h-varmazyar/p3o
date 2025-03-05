package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Create(ctx *gin.Context) {
	createLinkReq := new(CreateLinkReq)

	if err := ctx.ShouldBindJSON(&createLinkReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, ErrInvalidData.AddOriginalError(err), false)
		return
	}

	if !isValidLink(createLinkReq.RealUrl) {
		utils.JsonHttpResponse(ctx, nil, ErrInvalidLink, false)
		return
	}

	if createLinkReq.Key == "" {
		var err error
		createLinkReq.Key, err = pickKey()
		if err != nil {
			utils.JsonHttpResponse(ctx, nil, ErrKeyGenerationFailed, false)
			return
		}
	}

	linkData := &entities.Link{
		Key:       createLinkReq.Key,
		RealLink:  createLinkReq.RealUrl,
		Immediate: true,
	}

	if err := c.linkModel.Create(ctx, linkData); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	resp := &CreateLinkResp{
		Url:       linkData.RealLink,
		Key:       linkData.Key,
		Immediate: linkData.Immediate,
	}
	utils.JsonHttpResponse(ctx, resp, nil, true)
}