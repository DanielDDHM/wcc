package routes

import (
	"github.com/DanielDDHM/world-coin-converter/handlers/currencies"
	customratios "github.com/DanielDDHM/world-coin-converter/handlers/custom_ratios"
	"github.com/DanielDDHM/world-coin-converter/handlers/partners"
	"github.com/DanielDDHM/world-coin-converter/handlers/quote_categories"
	"github.com/DanielDDHM/world-coin-converter/handlers/quotes"
	"github.com/DanielDDHM/world-coin-converter/routes/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		partners.Routes(main.Group("partners", middlewares.Secret()))
		currencies.Routes(main.Group("currencies"))
		quotes.Routes(main.Group("quotes", middlewares.Auth()))
		quote_categories.Routes(main.Group("categories", middlewares.Secret()))
		customratios.Routes(main.Group("ratios"))
	}
	return router
}
