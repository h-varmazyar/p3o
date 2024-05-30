package utils

import (
	"github.com/gin-gonic/gin"
	localErr "github.com/h-varmazyar/p3o/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type httpResponse struct {
	Message interface{}
	Success bool
	Error   interface{}
}

func JsonHttpResponse(ctx *gin.Context, response interface{}, err error, success bool) {
	resp := new(httpResponse)

	statusCode := http.StatusOK

	if success {
		resp.Success = true
		resp.Message = response
	} else {
		castedErr := localErr.Cast(err)
		resp.Error = castedErr
		statusCode = castedErr.HttpCode
		log.WithError(err).Error(castedErr, castedErr.Details(), castedErr.Original())
	}

	ctx.IndentedJSON(statusCode, resp)
}
