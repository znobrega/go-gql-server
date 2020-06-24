package models

// User defines a user for the app
type Video struct {
	Title string  `gorm:"not null;unique_index:idx_email"`
	ID    *string // External user ID
	Url   *string
}
