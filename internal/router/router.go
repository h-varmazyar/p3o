package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/router/middlewares"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type linkService interface {
	ReturnByKey(ctx context.Context, key string) (domain.Link, error)
}

type Router struct {
	log      *log.Logger
	v1Router v1.Router
	linkSrv  linkService
}

func New(log *log.Logger, v1Router v1.Router, linkSrv linkService) Router {
	return Router{
		v1Router: v1Router,
		linkSrv:  linkSrv,
		log:      log,
	}
}

func (r Router) StartServing(ginEngine *gin.Engine, address string) error {
	handleRedirects(ginEngine, r.linkSrv)

	// ginEngine.Use(middlewares.CORSMiddleware())
	r.RegisterRoutes(ginEngine)

	srv := &http.Server{
		Addr:    address,
		Handler: ginEngine,
	}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
		return err
	}
	go srv.Serve(ln)

	return err
}

func (r Router) RegisterRoutes(ginRouter *gin.Engine) {
	r.log.Infof("********** registering routes")
	//ginRouter.Use(func(c *gin.Context) {
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
	//
	//	if c.Request.Method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//		return
	//	}
	//
	//	c.Next()
	//})
	apiRouter := ginRouter.Group("/api")
	r.v1Router.RegisterRoutes(apiRouter)
}

func handleRedirects(router *gin.Engine, linkSrv linkService) {
	router.GET("/:key", func(c *gin.Context) {
		key := c.Param("key")
		link, err := linkSrv.ReturnByKey(c, key)
		if err != nil {
			utils.JsonHttpResponse(c, nil, err, false)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, link.Url)
	})
}

func getIpAddress(c *gin.Context) string {
	return c.Request.RemoteAddr
}
