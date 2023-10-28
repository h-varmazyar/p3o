package link

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	linkGroup := router.Group("/link")
	linkGroup.POST("", c.createLink)
}
