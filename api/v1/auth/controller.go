package auth

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	log *log.Logger
}

func NewController(log *log.Logger) *Controller {
	return &Controller{
		log: log,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")

}
