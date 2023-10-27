package quote_categories

import (
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	database := database.GetDatabase()

	controller := &QuoteCategoryController{
		db: database,
	}

	g.POST("/", controller.NewQuoteCategory)

}
