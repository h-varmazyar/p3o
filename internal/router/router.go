package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	ginRouter *gin.RouterGroup
	v1Router  *v1.Router
}

func NewRouter(log *log.Logger) *Router {

}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	apiRouter := ginRouter.Group("/api")
	
	r.v1Router.RegisterRoutes(apiRouter)
}
