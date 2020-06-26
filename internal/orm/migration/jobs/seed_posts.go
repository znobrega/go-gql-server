package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/znobrega/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	firstPost *models.Post = &models.Post{
		ID:      "id2",
		Title:   "This is the first POST!2",
		Url:     "url test2",
		Comment: "test2",
	}
)

var SeedPosts *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_POSTS2",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstPost).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstPost).Error
	},
}
