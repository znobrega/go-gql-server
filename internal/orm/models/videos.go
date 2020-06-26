package models

type Video struct {
	Title string
	ID    *string
	Url   *string `gorm:"not null"`
}
