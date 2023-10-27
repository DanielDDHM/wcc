package partners

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DanielDDHM/world-coin-converter/cache"
	"github.com/DanielDDHM/world-coin-converter/handlers"
	"github.com/DanielDDHM/world-coin-converter/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PartnerController struct {
	db *gorm.DB
}

// NewPartner godoc
// @Summary      New Partner
// @Description  Create a new partner to consuming our api
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        partner body dtos.CreatePartnerDto true "Add Partner"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /partners [post]
func (s *PartnerController) NewPartner(c *gin.Context) {

	var p models.Partner

	fmt.Println(p)

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	err = p.Prepare()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var exist bool
	s.db.Model(&models.Partner{}).
		Select("count(*) > 0").Where("name = ? AND active = 1", p.Name).Scan(&exist)

	if exist {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Partner Already Exist")))
		return
	}

	err = s.db.Create(&p).Error
	p.PartnerId = ""
	p.PartnerKey = ""

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}

// GetPartner godoc
// @Summary      Get All Partner
// @Description  Get all partner
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /partners [get]
func (s *PartnerController) ListAll(c *gin.Context) {
	var p []models.Partner

	err := s.db.Select("id, name, country, locale, ttl, created_at, updated_at").Find(&p, "active = true").Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}

// GetPartner godoc
// @Summary      Revogate Partner
// @Description  Revogate partner to use our api
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Partner id"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /partners/disable/{id} [put]
func (s *PartnerController) DisablePartner(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var p models.Partner

	err = s.db.Find(&p, "id = ? AND active = true", newId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	if p.ID == 0 {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("partner not found")))
		return
	}

	p.Active = 0

	err = s.db.Save(&p).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	cacheKey := "partnermidd" + p.PartnerKey + p.PartnerId
	cache.Del(cacheKey)
	c.JSON(http.StatusOK, handlers.SuccessResponse(p))
}
