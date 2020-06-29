package resolvers

import (
	"context"

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

func (r *queryResolver) Orders(ctx context.Context, id *string, limit *int, page *int) (*models.Orders, error) {
	whereID := "id = ?"
	var dbRecords []*models.Order

	db := r.ORM.DB.New()

	if id == nil {
		db = db.Where(whereID, nil)
	}

	if id != nil && *id != "-99999" {
		db = db.Where(whereID, *id)
	}

	pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    *page,
		Limit:   *limit,
		OrderBy: []string{"id asc"},
	}, &dbRecords)

	count := len(dbRecords)

	response := &models.Orders{
		Limit: limit,
		Page:  page,
		Count: &count,
		List:  dbRecords,
	}

	return response, db.Error
}
