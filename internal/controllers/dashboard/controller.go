package dashboard

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

type linkService interface {
	List(ctx context.Context, userId uint) (domain.LinkList, error)
	TotalLinkCount(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
}

type visitService interface {
	TodayInfo(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	TotalVisit(ctx context.Context, userId uint) (domain.DashboardInfoItem, error)
	VisitChart(ctx context.Context, userId uint) ([]domain.ChartItem, error)
}

type Controller struct {
	linkSrv  linkService
	visitSrv visitService
}

func New(linkSrv linkService, visitSrv visitService) Controller {
	return Controller{
		linkSrv:  linkSrv,
		visitSrv: visitSrv,
	}
}
