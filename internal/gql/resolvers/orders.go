package resolvers

import (
	"context"
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/znobrega/go-gql-server/internal/gql/models"
)

func (r *queryResolver) Order(ctx context.Context) (*models.Order, error) {
	res := &models.Order{
		ID:           2,
		CustomerName: "asdasdasd",
		OrderAmount:  2,
	}

	return res, nil
}

func (r *orderResolver) OrderAmount(ctx context.Context, obj *models.Order) (float64, error) {
	panic("not implemented")
}

func (r *queryResolver) Orders(ctx context.Context, id *string, limit *int, page *int) ([]*models.Order, error) {
	whereID := "id = ?"
	var dbRecords []*models.Order
	db := r.ORM.DB.New()

	if id == nil {
		db = db.Where(whereID, nil)
	}

	if id != nil && *id != "-99999" {
		db = db.Where(whereID, *id)
	}

	fmt.Println("test")

	pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    *page,
		Limit:   *limit,
		OrderBy: []string{"id asc"},
	}, &dbRecords)

	return dbRecords, db.Error
}
