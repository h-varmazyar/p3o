package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://127.0.0.1",
			"http://localhost",
			"http://localhost:6854",
			"https://p3o.ir",
			"https://v0-p3o-ui.vercel.app",
			"https://v0-p3o-ui.vercel.app/",
		}

		allowedPatterns := []string{
			"*.p3o.ir",
		}

		origin := c.Request.Header.Get("Origin")
		isAllowed := false

		for _, o := range allowedOrigins {
			if origin == o {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			for _, pattern := range allowedPatterns {
				domain := strings.TrimPrefix(pattern, "*.")
				if strings.HasSuffix(origin, domain) {
					isAllowed = true
					break
				}
			}
		}

		fmt.Println("cors origin is:", origin)

		if isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		if c.Request.Method == http.MethodOptions {
			if isAllowed {
				c.AbortWithStatus(http.StatusNoContent)
			} else {
				c.AbortWithStatus(http.StatusForbidden)
			}
			return
		}

		c.Next()
	}
}
