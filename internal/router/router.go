package router

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Router struct {
	v1Router *v1.Router
	log      *log.Logger
}

type Params struct {
	fx.In

	Log       *log.Logger
	GinEngine *gin.Engine
	V1Router  *v1.Router
}

type Result struct {
	fx.Out

	Router *Router
}

func New(lc fx.Lifecycle, params Params) Result {
	router := &Router{
		v1Router: params.V1Router,
		log:      params.Log,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.RegisterRoutes(params.GinEngine)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return Result{Router: router}
}

func (r *Router) RegisterRoutes(ginRouter *gin.Engine) {
	apiRouter := ginRouter.Group("/api")

	r.log.Infof("********** stating router")

	r.v1Router.RegisterRoutes(apiRouter)
}
