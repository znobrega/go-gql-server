package resolvers

import (
	"context"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/iancoleman/strcase"
	"github.com/znobrega/go-gql-server/internal/gql/models"
)

func (r *queryResolver) Order(ctx context.Context) (*models.Order, error) {
	res := &models.Order{
		ID:           2,
		CustomerName: "asdasdasd",
		Amount:       2,
	}

	return res, nil
}

func (r *queryResolver) OrderAmount(ctx context.Context, obj *models.Order) (float64, error) {
	panic("not implemented")
}

func (r *queryResolver) Orders(ctx context.Context, limit *int, page *int, filter map[string]interface{}) (*models.Orders, error) {
	var dbRecords []*models.Order

	db := r.ORM.DB.New()

	if filter != nil {
		filterSnakeCase := make(map[string]interface{})

		for key, _ := range filter {
			filterSnakeCase[strcase.ToSnake(key)] = filter[key]
		}

		db = db.Where(filterSnakeCase)
	}

	pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    *page,
		Limit:   *limit,
		OrderBy: []string{"id desc"},
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
