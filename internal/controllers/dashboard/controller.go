package dashboard

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/pagination"
)

type linkService interface {
	List(ctx context.Context, userId uint, paginator pagination.Paginator) (domain.LinkList, error)
	TotalLinkCount(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	TodayInfo(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	TotalVisit(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	VisitChart(ctx context.Context, userId uint, key string) ([]domain.ChartItem, error)
}

type Controller struct {
	linkSrv linkService
}

func New(linkSrv linkService) Controller {
	return Controller{
		linkSrv: linkSrv,
	}
}
