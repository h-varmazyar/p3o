package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/panel"
	"github.com/h-varmazyar/p3o/internal/controllers/visit"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type Controller struct {
	log   *log.Logger
	panel *panel.Controller
}

func NewController(log *log.Logger, linkService linkService.Service) error {
	c := &Controller{
		log:   log,
		panel: panel.NewController(log, linkService),
	}

	groupErr := new(errgroup.Group)

	c.RegisterPanelServer(log, linkService, groupErr)
	c.RegisterVisitServer(log, linkService, groupErr)

	if err := groupErr.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *Controller) RegisterPanelServer(log *log.Logger, linkService linkService.Service, groupErr *errgroup.Group) {
	e := gin.New()
	e.Use(gin.Recovery())

	panelController := panel.NewController(log, linkService)
	panelController.RegisterRoutes(e)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	groupErr.Go(func() error {
		return server.ListenAndServe()
	})
}

func (c *Controller) RegisterVisitServer(log *log.Logger, linkService linkService.Service, groupErr *errgroup.Group) {
	e := gin.New()
	e.Use(gin.Recovery())

	visitController := visit.NewController(log, linkService)
	visitController.RegisterRoutes(e)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	groupErr.Go(func() error {
		return server.ListenAndServe()
	})
}
