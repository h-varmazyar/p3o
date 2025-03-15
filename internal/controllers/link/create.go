package link

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
=======
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
>>>>>>> 292128d (feat: add link creation)
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) Create(ctx *gin.Context) {
<<<<<<< HEAD
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
=======
	linkCreateReq := new(domain.LinkCreateReq)

	if err := ctx.ShouldBindJSON(&linkCreateReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, errors.ErrInvalidData.AddOriginalError(err), false)
		return
	}

	linkCreateReq.UserId=utils.FetchUserId(ctx)

	resp, err:=c.linkService.Create(ctx, *linkCreateReq)
	if err!=nil{
>>>>>>> 292128d (feat: add link creation)
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

<<<<<<< HEAD
	resp := &CreateLinkResp{
		Url:       linkData.RealLink,
		Key:       linkData.Key,
		Immediate: linkData.Immediate,
	}
=======
>>>>>>> 292128d (feat: add link creation)
	utils.JsonHttpResponse(ctx, resp, nil, true)
}