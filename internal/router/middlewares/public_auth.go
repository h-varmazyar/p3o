package middlewares

import (
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/jwt"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PublicAuthMiddleware struct {
	logger *log.Logger
}

func NewPublicAuthMiddleware(logger *log.Logger) PublicAuthMiddleware {
	return PublicAuthMiddleware{
		logger: logger,
	}
}

func (p PublicAuthMiddleware) respondUnauthorized(c *gin.Context) {
	utils.JsonHttpResponse(c, nil, errors.ErrUnauthorized, false)
	c.Abort()
}

func (p PublicAuthMiddleware) Handle(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		p.respondUnauthorized(c)
		return
	}
	authToken = strings.TrimSpace(strings.TrimPrefix(authToken, "Bearer"))
	claim, _, err := jwt.GetClaim(authToken)
	if err != nil {
		p.respondUnauthorized(c)
		return
	}

	userID, err := strconv.Atoi(claim)
	if err != nil {
		p.respondUnauthorized(c)
		return
	}

	uID := uint(userID)
	c.Set("user_id", uID)
	c.Next()
}
