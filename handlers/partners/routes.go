package partners

import (
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	database := database.GetDatabase()

	controller := &PartnerController{
		db: database,
	}

	g.POST("/", controller.NewPartner)
	g.GET("/", controller.ListAll)
	g.PUT("/disable/:id", controller.DisablePartner)

}
