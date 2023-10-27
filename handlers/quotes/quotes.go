package quotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DanielDDHM/world-coin-converter/dtos"
	"github.com/DanielDDHM/world-coin-converter/handlers"
	"github.com/DanielDDHM/world-coin-converter/integrations/ows"
	"github.com/DanielDDHM/world-coin-converter/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuoteController struct {
	db *gorm.DB
}

// NewQuote godoc
// @Summary      New Quote
// @Description  Create a new quote
// @Tags         quotes
// @Accept       json
// @Produce      json
// @Param        partner body dtos.RequestQuoteDto true "Add Quote"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /quotes [post]
func (s *QuoteController) NewQuote(c *gin.Context) {

	partner, _ := c.Get("partner")

	partnerId, ok := partner.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Failed to parse partner"))
		return
	}

	var pr models.Partner

	s.db.Select("id, ttl").First(&pr, "id = ?", partnerId)

	var p dtos.RequestQuoteDto
	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	validator := validator.New()
	err = validator.Struct(p)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var currencyDestination models.Currency

	err = s.db.Select("id, iso_code").
		First(&currencyDestination, "iso_code = ?", p.DestinationCurrency).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	var currencyOrigin models.Currency
	err = s.db.Select("id, iso_code").
		First(&currencyOrigin, "iso_code = ?", p.OriginCurrency).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	// var cr models.CustomRatio

	// s.db.First(&cr, "target = ? AND currency = ?", p.UserIdentifier, currencyDestination.ID)
	var side string
	// var spread float64

	if currencyOrigin.Type == "crypto" {
		side = "sell"
	} else {
		side = "buy"
	}

	data := dtos.RequestQuoteOwsDto{
		Trading:    p.OriginCurrency,
		Settlement: p.DestinationCurrency,
		Side:       side,
		Size:       p.Size,
		Amount:     p.Amount,
	}

	res, err := json.Marshal(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	result, err := ows.OwsRequest("POST", "quote", string(res))

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	partnerResult, err := json.Marshal(result)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	// Moeda que a plataforma tá recebendo é o BID
	// Moeda que a platagorma tá dando é o ASK
	// err = s.db.Find(&qc, "name = ? AND currency = ?", p.UserIdentifier).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	// Price:                 fmt.Sprint(result["price"].(float64) * 0.1), => 0.1 será trocapo pelo spread

	// Aqui será o spread
	// a := "0.8"

	// a -> Estou vendendo - Te oferecendo para colocar o que tenho - BID
	// b -> Estou comprando -> Quanto que a plataforma está pedindo - ASK

	// b, err := strconv.ParseFloat(a, 64)

	quote := models.Quote{
		PartnerId:             partnerId,
		Price:                 fmt.Sprint(result["price"].(float64)),
		CurrencyOriginId:      currencyOrigin.ID,
		CurrencyDestinationId: currencyDestination.ID,
		PartnerResult:         partnerResult,
		ExpiresIn:             time.Now().Add(time.Second * time.Duration(pr.Ttl)),
		Size:                  fmt.Sprint(result["size"].(float64)),
		QuoteId:               uuid.New().String(),
		UserIdentifier:        p.UserIdentifier,
	}

	err = quote.Prepare()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	err = s.db.Create(&quote).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, handlers.SuccessResponse(quote))
}

// ConfirmQuote godoc
// @Summary      Confirm Quote
// @Description  Confirm a quote
// @Tags         quotes
// @Accept       json
// @Produce      json
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.HTTPError
// @Failure      400  {object}  handlers.HTTPError
// @Failure      401  {object}  handlers.HTTPUnauthorized
// @Router       /quotes/confirm/{id} [put]
func (s *QuoteController) ConfirmQuote(c *gin.Context) {
	partner, _ := c.Get("partner")

	quoteId := c.Param("id")
	partnerId, ok := partner.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Failed to parse partner"))
		return
	}

	var q models.Quote
	err := s.db.First(&q, "quote_id = ? AND partner_id = ?", quoteId, partnerId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(err))
		return
	}

	if time.Now().After(q.ExpiresIn) {
		q.Status = "expired"
		s.db.Save(&q)
		c.JSON(http.StatusBadRequest, handlers.ErrorResponse(errors.New("Quote expires")))
		return
	}

	q.Status = "confirmed"

	s.db.Save(&q)

	c.JSON(http.StatusOK, handlers.SuccessResponse(q))
}
