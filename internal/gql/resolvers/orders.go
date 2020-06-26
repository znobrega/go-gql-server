package resolvers

import (
	"context"

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

func (r *queryResolver) Orders(ctx context.Context, id *string) ([]*models.Order, error) {
	return orderList(r)
}

func orderList(r *queryResolver) ([]*models.Order, error) {
	var id *string
	whereID := "id = ?"
	var dbRecords []*models.Order
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords)

	return dbRecords, db.Error
}
