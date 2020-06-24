package resolvers

import (
	"context"

	log "github.com/znobrega/go-gql-server/internal/logger"

	"github.com/znobrega/go-gql-server/internal/gql/models"
	tf "github.com/znobrega/go-gql-server/internal/gql/resolvers/transformations"
	dbm "github.com/znobrega/go-gql-server/internal/orm/models"
)

func (r *queryResolver) Videos(ctx context.Context, id *string) (*models.Videos, error) {
	return videosList(r, id)
}

func videosList(r *queryResolver, id *string) (*models.Videos, error) {
	entity := "users"
	whereID := "id = ?"
	record := &models.Videos{}
	dbRecords := []*dbm.Video{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBVideoToGQLVideo(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
