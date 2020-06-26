package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/znobrega/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                    = "Test User"
	fname                    = "Test"
	lname                    = "User"
	nname                    = "Foo Bar"
	description              = "This is the first user ever!"
	location                 = "His house, maybe? Wouldn't know"
	firstUser   *models.User = &models.User{
		Email:       "test@test.com2sadsasd",
		Name:        &uname,
		FirstName:   &fname,
		LastName:    &lname,
		NickName:    &nname,
		Description: &description,
		Location:    &location,
	}
)

var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS2",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
