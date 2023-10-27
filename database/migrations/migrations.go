package migrations

import (
	"github.com/DanielDDHM/world-coin-converter/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {

	db.AutoMigrate(
		models.Partner{},
		models.Currency{},
		models.Quote{},
		models.QuoteCategory{},
		models.CustomRatio{},
	)
}
