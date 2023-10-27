package middlewares

import (
	"net/http"

	"github.com/DanielDDHM/world-coin-converter/config"
	"github.com/gin-gonic/gin"
)

func Secret() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("x-api-secret")

		if secret == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if config.GetConfig().ApiSecret != secret {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
