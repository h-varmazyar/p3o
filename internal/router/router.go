package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/entities"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	"github.com/h-varmazyar/p3o/internal/workers"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net"
	"net/http"
)

type linkService interface {
	ReturnByKey(ctx context.Context, key string) (entities.Link, error)
}

type Router struct {
	v1Router *v1.Router
	log      *log.Logger
}

type Params struct {
	fx.In

	Log       *log.Logger
	GinEngine *gin.Engine
	V1Router  *v1.Router
	VisitChan chan workers.VisitRecord
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
	handleRedirects(params.GinEngine, nil, params.VisitChan)
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
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	apiRouter := ginRouter.Group("/api")
	r.v1Router.RegisterRoutes(apiRouter)
}

func handleRedirects(router *gin.Engine, linkModel linkService, visitChannel chan workers.VisitRecord) {
	router.GET("/:key", func(c *gin.Context) {
		key := c.Param("key")
		link, err := linkModel.ReturnByKey(c, key)
		if err != nil {
			utils.JsonHttpResponse(c, nil, err, false)
			return
		}
		visitChannel <- workers.VisitRecord{
			LinkId:    link.ID,
			IpAddress: getIpAddress(c),
		}
		c.Redirect(http.StatusTemporaryRedirect, link.RealLink)
	})
}

func getIpAddress(c *gin.Context) string {
	return c.Request.RemoteAddr
}
