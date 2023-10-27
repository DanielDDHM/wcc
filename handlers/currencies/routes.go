package currencies

import (
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	database := database.GetDatabase()

	controller := &CurrencyController{
		db: database,
	}

	g.POST("/", controller.NewCurrency)
	g.GET("/", controller.ListAll)

}
