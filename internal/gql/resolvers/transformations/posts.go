package transformations

import (
	gql "github.com/znobrega/go-gql-server/internal/gql/models"
	dbm "github.com/znobrega/go-gql-server/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBPostToGQLPost(i *dbm.Post) (o *gql.Post, err error) {
	o = &gql.Post{
		ID:      i.ID,
		Title:   i.Title,
		URL:     i.Url,
		Comment: i.Comment,
	}
	return o, err
}
