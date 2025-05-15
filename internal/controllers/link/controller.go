package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/pagination"
)

type linkService interface {
	Create(ctx context.Context, link domain.LinkCreateReq) (domain.LinkCreateResp, error)
	List(ctx context.Context, userId uint, paginator pagination.Paginator) (domain.LinkList, error)
	Edit(ctx context.Context, req domain.EditLinkReq) error
	Activate(ctx context.Context, userId uint, key string) error
	Deactivate(ctx context.Context, userId uint, key string) error
	TotalLinkCount(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	Status(ctx context.Context, userId uint, key string) (string, error)
	Fetch(ctx context.Context, userId uint, key string) (domain.LinkDetails, error)
	Delete(ctx context.Context, userId uint, key string) error
	IndirectVisit(ctx context.Context, key string, id string) (domain.Link, error)
}

type Controller struct {
	linkService linkService
}

func New(linkSrv linkService) Controller {
	return Controller{
		linkService: linkSrv,
	}
}
