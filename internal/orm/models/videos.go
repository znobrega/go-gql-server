package models

// User defines a user for the app
type Video struct {
	Title string
	ID    *string
	Url   *string `gorm:"not null"`
}
