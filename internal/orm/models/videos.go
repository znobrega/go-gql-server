package models

// User defines a user for the app
type Video struct {
	BaseModelSoftDelete         // We don't to actually delete the users, maybe audit if we want to hard delete them? or wait x days to purge from the table, also
	Title               string  `gorm:"not null;unique_index:idx_email"`
	Id                  *string // External user ID
	Url                 *string
}
