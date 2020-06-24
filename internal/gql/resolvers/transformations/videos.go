package transformations

import (
	gql "github.com/znobrega/go-gql-server/internal/gql/models"
	dbm "github.com/znobrega/go-gql-server/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBVideoToGQLVideo(i *dbm.Video) (o *gql.Video, err error) {
	o = &gql.Video{
		ID:    *i.ID,
		Title: i.Title,
		URL:   *i.Url,
	}
	return o, err
}
