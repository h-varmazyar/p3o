package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net"
	"net/http"
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
	router.RegisterRoutes(params.GinEngine)

	srv := &http.Server{
		Addr:    ":8765",
		Handler: params.GinEngine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
				return err
			}
			go srv.Serve(ln)
			fmt.Println("[My Demo]Succeeded to start HTTP Server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return Result{Router: router}
}

func (r *Router) RegisterRoutes(ginRouter *gin.Engine) {
	r.log.Infof("********** registering routes")
	apiRouter := ginRouter.Group("/api")
	r.v1Router.RegisterRoutes(apiRouter)
}
