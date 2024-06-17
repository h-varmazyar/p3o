package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/entities"
	linkModel "github.com/h-varmazyar/p3o/internal/models/link"
	"github.com/h-varmazyar/p3o/internal/workers"
	"github.com/h-varmazyar/p3o/pkg/environments"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"go.uber.org/fx"
)

var (
	configs *config
)

func init() {
	configs = new(config)
	err := environments.LoadEnvironments(configs)
	if err != nil {
		panic(fmt.Sprintf("failed to load auth controller configs: %v", err))
	}
}

type Controller struct {
	linkModel linkModel.Model
	//LinkCache linkModel.Model
	VisitChan chan workers.VisitRecord
}

type Params struct {
	fx.In

	LinkModel linkModel.Model
	LinkCache linkModel.Model
	VisitChan chan workers.VisitRecord
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{
		linkModel: p.LinkModel,
		//LinkCache: p.LinkCache,
		VisitChan: p.VisitChan,
	}
	return Result{Controller: controller}
}

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

//func (c *Controller) Fetch(ctx *gin.Context) {
//	req := new(FetchLinkReq)
//	link, err := c.LinkCache.ReturnByKey(ctx, req.Key)
//	if err != nil && errors.As(err, &linkModel.ErrCacheFetchFailed) {
//		utils.JsonHttpResponse(ctx, nil, err, false)
//		return
//	}
//
//	if link == nil {
//		link, err = c.linkModel.ReturnByKey(ctx, req.Key)
//		if err != nil {
//			utils.JsonHttpResponse(ctx, nil, err, false)
//			return
//		}
//
//		if err = c.LinkCache.Create(ctx, link); err != nil {
//			utils.JsonHttpResponse(ctx, nil, err, false)
//			return
//		}
//	}
//
//	resp := &Link{
//		Url:       link.RealLink,
//		Immediate: link.Immediate,
//	}
//
//	c.VisitChan <- workers.VisitRecord{LinkId: link.ID}
//
//	utils.JsonHttpResponse(ctx, resp, nil, true)
//}
