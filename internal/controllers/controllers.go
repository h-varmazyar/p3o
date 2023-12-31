package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/panel"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	log   *log.Logger
	panel *panel.Controller
}

func NewController(log *log.Logger, linkService linkService.Service) *Controller {
	c := &Controller{
		log:   log,
		panel: panel.NewController(log, linkService),
	}
	return c
}

func (c *Controller) RegisterRoutes(ginEngine *gin.Engine) {
	c.panel.RegisterRoutes(&ginEngine.RouterGroup)
}
