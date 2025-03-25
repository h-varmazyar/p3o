package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/workers"
)

type linkService interface {
	Create(ctx context.Context, link domain.LinkCreateReq) (domain.LinkCreateResp, error)
	All(ctx context.Context, userId uint) ([]entities.Link, error)
	Activate(ctx context.Context, userId uint, key string) error
	Deactivate(ctx context.Context, userId uint, key string) error
	TotalVisits(ctx context.Context, userId uint) (uint, error)
	TotalLinkCount(ctx context.Context, userId uint) (uint, error)
	Status(ctx context.Context, userId uint, key string) (string, error)
	Delete(ctx context.Context, userId uint, key string) error
}

type Controller struct {
	VisitChan   chan workers.VisitRecord
	linkService linkService
}

func New(linkSrv linkService, visitChan chan workers.VisitRecord) Controller {
	return Controller{
		VisitChan:   visitChan,
		linkService: linkSrv,
	}
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
