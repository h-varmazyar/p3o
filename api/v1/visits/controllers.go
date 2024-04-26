package visit

import (
	"github.com/gin-gonic/gin"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	visitsRouter := router.Group("/visits")

}
