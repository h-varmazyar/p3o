package link

import (
	"github.com/gin-gonic/gin"
	linkService "github.com/h-varmazyar/p3o/internal/repositories/link/service"
	"github.com/h-varmazyar/p3o/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	log         *log.Logger
	linkService linkService.Service
}

func NewController(log *log.Logger, linkService linkService.Service) *Controller {
	return &Controller{
		log:         log,
		linkService: linkService,
	}
}

func (c *Controller) createLink(ctx *gin.Context) {
	req := new(linkService.CreateLinkReq)
	ctx.ShouldBindJSON(&req)

	link, err := c.linkService.CreateLink(ctx, req)
	if err != nil {
		e := errors.Cast(err)
		ctx.JSON(e.HttpCode, e.Json(ctx))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"link": link.Url,
	})
}
