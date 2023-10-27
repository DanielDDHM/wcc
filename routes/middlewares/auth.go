package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/DanielDDHM/world-coin-converter/cache"
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/DanielDDHM/world-coin-converter/models"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		partner_key := c.GetHeader("x-client-key")
		partner_id := c.GetHeader("x-client-id")

		if partner_key == "" || partner_id == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		cacheKey := "partnermidd" + partner_key + partner_id
		cached, _ := cache.Get(cacheKey)

		if cached != nil {
			data := models.Partner{}

			bytes, _ := json.Marshal(cached)
			json.Unmarshal(bytes, &data)

			c.Set("partner", data.ID)
			c.Next()
			return
		}

		db := database.GetDatabase()

		var p models.Partner

		err := db.Select("id, partner_key, partner_id, active").First(&p, `partner_key = ? AND partner_id = ?
		AND active = true`, partner_key, partner_id).Error

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		json, _ := json.Marshal(p)

		cache.Set(cacheKey, string(json), (60*60)*24)
		c.Set("partner", p.ID)
		c.Next()
	}
}
