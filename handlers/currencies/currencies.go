package currencies

import (
	"errors"
	"net/http"

	"github.com/DanielDDHM/world-coin-converter/handlers"
	"github.com/DanielDDHM/world-coin-converter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CurrencyController struct {
	db *gorm.DB
}

// NewCurrency godoc
// @Summary      New Currency
// @Description  Create a new currency
// @Tags         currencies
// @Accept       json
// @Produce      json
// @Param        partner body dtos.CurrencyDto true "Add Currency"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Router       /currencies [post]
func (s *CurrencyController) NewCurrency(c *gin.Context) {
	var p models.Currency
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var exist bool
	s.db.Model(&models.Currency{}).
		Select("count(*) > 0").Where("name = ? OR iso_code = ?", p.Name, p.IsoCode).Scan(&exist)

	if exist {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Currency already exist")))
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	err = p.Prepare()
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	err = s.db.Create(&p).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}

// GetCurrencies godoc
// @Summary      Get All Currencies
// @Description  Get all currencies
// @Tags         currencies
// @Accept       json
// @Produce      json
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /currencies [get]
func (s *CurrencyController) ListAll(c *gin.Context) {
	var p []models.Currency

	err := s.db.Find(&p, "active = true").Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}
