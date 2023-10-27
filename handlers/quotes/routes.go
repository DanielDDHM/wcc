package quotes

import (
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	database := database.GetDatabase()

	controller := &QuoteController{
		db: database,
	}

	// g.POST("/", controller.NewPartner)
	g.POST("/", controller.NewQuote)
	g.PUT("/confirm/:id", controller.ConfirmQuote)

}
