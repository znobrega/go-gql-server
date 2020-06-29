package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/znobrega/go-gql-server/internal/gql/models"
	"gopkg.in/gormigrate.v1"
)

var (
	firstOrder *models.Order = &models.Order{
		ID:           2,
		CustomerName: "NILBSON GALINDO",
		Amount:       3,
	}
)

var SeedOrder *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_ORDER",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstOrder).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstOrder).Error
	},
}
