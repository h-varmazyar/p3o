package link

import (
	"context"
	"fmt"

	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/workers"
	"github.com/h-varmazyar/p3o/pkg/environments"
	"go.uber.org/fx"
)

var (
	configs *config
)

//todo: must be deleted
func init() {
	configs = new(config)
	err := environments.LoadEnvironments(configs)
	if err != nil {
		panic(fmt.Sprintf("failed to load auth controller configs: %v", err))
	}
}

type linkService interface{
	Create(ctx context.Context, link domain.LinkCreateReq) (domain.LinkCreateResp, error)
	All(ctx context.Context, userId uint) ([]entities.Link, error)
	Activate(ctx context.Context, userId uint, key string) error
	Deactivate(ctx context.Context, userId uint, key string) error
	TotalVisits(ctx context.Context, userId uint) (uint, error)
	TotalLinkCount(ctx context.Context, userId uint) (uint, error)
	Status(ctx context.Context,userId uint, key string) (string, error)
	Delete(ctx context.Context, userId uint, key string) error
}

type Controller struct {
	VisitChan chan workers.VisitRecord
	linkService linkService
}

type Params struct {
	fx.In

	VisitChan chan workers.VisitRecord
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{
		//LinkCache: p.LinkCache,
		VisitChan: p.VisitChan,
	}
	return Result{Controller: controller}
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
