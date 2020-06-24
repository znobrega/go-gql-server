package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/znobrega/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	id                       = "123"
	url                      = "https://VIDEO.COM"
	firstVideo *models.Video = &models.Video{
		Title: "This is the first VIDEO!",
		Url:   &url,
	}
)

// SeedUsers inserts the first users
var SeedVideos *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_VIDEOS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstVideo).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstVideo).Error
	},
}
