package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/znobrega/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	id                       = "1232"
	url                      = "https://VIDEO.COM2"
	firstVideo *models.Video = &models.Video{
		Title: "This is the first VIDEO!2",
		Url:   &url,
	}
)

// SeedUsers inserts the first users
var SeedVideos *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_VIDEOS2",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstVideo).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstVideo).Error
	},
}
