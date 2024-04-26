package links

import (
	"github.com/gin-gonic/gin"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
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

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	linksRouter := router.Group("/links")
	linksRouter.POST("/", c.create)
	linksRouter.GET("/:key", c.fetch)
}

func (c *Controller) create(ctx *gin.Context) {
	link := new(Link)
	if err := ctx.ShouldBindJSON(link); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	if newLink, err := c.linkService.CreateLink(ctx, &linkService.CreateLinkReq{RealUrl: link.URL, Key: link.Key}); err != nil {
		e := errors.Cast(err)
		ctx.JSON(e.HttpCode, e.Json(ctx))
		return
	} else {
		ctx.JSON(http.StatusCreated, newLink)
	}
}

func (c *Controller) fetch(ctx *gin.Context) {
	key := new(Key)
	if err := ctx.ShouldBindUri(key); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if link, err := c.linkService.FetchLink(ctx, &linkService.FetchLinkReq{Key: key.Key}); err != nil {
		e := errors.Cast(err)
		ctx.JSON(e.HttpCode, e.Json(ctx))
		return
	} else {
		if link.Immediate {
			ctx.Redirect(http.StatusPermanentRedirect, link.Url)
			return
		} else {
			ctx.Redirect(http.StatusTemporaryRedirect, link.Url)
		}
	}
}
