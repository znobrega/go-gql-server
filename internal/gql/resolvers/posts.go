package resolvers

import (
	"context"

	log "github.com/znobrega/go-gql-server/internal/logger"

	"github.com/znobrega/go-gql-server/internal/gql/models"
	tf "github.com/znobrega/go-gql-server/internal/gql/resolvers/transformations"
	dbm "github.com/znobrega/go-gql-server/internal/orm/models"
)

func (r *queryResolver) Posts(ctx context.Context, id *string) (*models.Posts, error) {
	return postList(r, id)
}

func postList(r *queryResolver, id *string) (*models.Posts, error) {
	entity := "posts"
	whereID := "id = ?"
	record := &models.Posts{}
	dbRecords := []*dbm.Post{}
	db := r.ORM.DB.New()

	if id != nil {
		db = db.Where(whereID, *id)
	}

	db = db.Find(&dbRecords).Count(&record.Count)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBPostToGQLPost(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
