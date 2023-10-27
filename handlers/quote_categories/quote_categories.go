package quote_categories

import (
	"errors"
	"net/http"

	"github.com/DanielDDHM/world-coin-converter/handlers"
	"github.com/DanielDDHM/world-coin-converter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuoteCategoryController struct {
	db *gorm.DB
}

func (s *QuoteCategoryController) NewQuoteCategory(c *gin.Context) {
	var p models.QuoteCategory

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var exist bool
	s.db.Model(&models.Partner{}).
		Select("count(*) > 0").
		Where("id = ? AND active = 1", p.PartnerId).Scan(&exist)

	if !exist {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Partner Not Exist")))
		return
	}

	var existCurrency bool
	s.db.Model(&models.Currency{}).
		Select("count(*) > 0").
		Where("id = ? AND active = 1", p.PartnerId).Scan(&existCurrency)

	if !existCurrency {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Currency Not Exist")))
		return
	}

	s.db.Find(&p, "name = ? AND partner_id = ?", p.Name, p.PartnerId)

	if p.ID != 0 {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Category Already Exist")))
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
