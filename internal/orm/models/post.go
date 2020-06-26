package models

type Post struct {
	ID      string
	Title   string
	Url     string `gorm:"not null;unique_index:idx_url"`
	Comment string
}
