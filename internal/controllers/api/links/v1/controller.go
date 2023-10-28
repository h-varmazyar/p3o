package link

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Controller interface {
}

type controller struct {
	log *log.Logger
}

func NewController(log *log.Logger, gin *gin.RouterGroup) Controller {
	return &controller{log: log}
}
