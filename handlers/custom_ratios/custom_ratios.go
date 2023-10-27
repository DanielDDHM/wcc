package customratios

import (
	"errors"
	"net/http"

	"github.com/DanielDDHM/world-coin-converter/dtos"
	"github.com/DanielDDHM/world-coin-converter/handlers"
	"github.com/DanielDDHM/world-coin-converter/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomRatiosController struct {
	db *gorm.DB
}

// NewCustomRatio godoc
// @Summary      New Custom Ratio
// @Description  Create a new custom ratio
// @Tags         Custom Ratio
// @Accept       json
// @Produce      json
// @Param        partner body dtos.CustomRatioDto true "Add Custom Ratio"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /ratios [post]
func (s *CustomRatiosController) NewCustomRatio(c *gin.Context) {

	var dto dtos.CustomRatioDto

	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	validator := validator.New()
	err = validator.Struct(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var partner models.Partner
	err = s.db.First(&partner, "id = ? AND active = 1", dto.PartnerId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Partner Not Found")))
		return
	}

	var currency models.Currency

	err = s.db.First(&currency, "iso_code = ?", dto.Currency).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var exist bool

	s.db.Model(&models.CustomRatio{}).
		Select("count(*) > 0").
		Where("target = ? AND currency_id = ? AND partner_id = ?", dto.Target, currency.ID, dto.PartnerId).Scan(&exist)

	if exist {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Custom Ratio Already exist")))
		return
	}

	p := &models.CustomRatio{
		Target:     dto.Target,
		CurrencyId: currency.ID,
		AskRatio:   dto.AskRatio,
		BuyRatio:   dto.BuyRatio,
		PartnerId:  dto.PartnerId,
	}

	p.Prepare()

	err = s.db.Create(&p).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	p.Currency = nil

	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}
