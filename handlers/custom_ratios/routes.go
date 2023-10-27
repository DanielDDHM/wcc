package customratios

import (
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	database := database.GetDatabase()

	controller := &CustomRatiosController{
		db: database,
	}

	g.POST("/", controller.NewCustomRatio)
}
