package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Engine
	v1 *v1.Controller
}

func NewController() *Controller {
	c := &Controller{
		Engine: gin.Default(),
	}
	return c
}

func (c *Controller) RegisterRoutes() {
	c.v1.RegisterRoutes()
}
