package resolvers

import (
	"context"
	"fmt"
	"log"

	"github.com/iancoleman/strcase"
	paginator "github.com/pilagod/gorm-cursor-paginator"
	"github.com/znobrega/go-gql-server/internal/gql/models"
)

type PagingQuery struct {
	After  *string
	Before *string
	Limit  *int
	Order  *string
}
type Cursor struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
}

func GetModelPaginator(q PagingQuery) *paginator.Paginator {
	p := paginator.New()

	if q.After != nil {
		p.SetAfterCursor(*q.After) // [default: nil]
	}

	if q.Before != nil {
		p.SetBeforeCursor(*q.Before) // [default: nil]
	}

	if q.Limit != nil {
		p.SetLimit(*q.Limit) // [default: 10]
	}

	if q.Order != nil && *q.Order == "asc" {
		p.SetOrder(paginator.ASC) // [default: paginator.DESC]
	}
	return p
}

func (r *queryResolver) Order(ctx context.Context) (*models.Order, error) {
	res := &models.Order{
		ID:           2,
		CustomerName: "asdasdasd",
		Amount:       2,
	}

	return res, nil
}

func (r *queryResolver) Orders(ctx context.Context, limit *int, page *int, filter map[string]interface{}) (*models.Orders, error) {
	var dbRecords []*models.Order

	db := r.ORM.DB.New()

	if filter != nil {
		fmt.Println(filter)
		filterSnakeCase := make(map[string]interface{})

		for key, val := range filter {
			fmt.Println(filter)
			filterSnakeCase[strcase.ToSnake(key)] = val
		}

		db = db.Where(filterSnakeCase)
	}
	fmt.Printf("Passou 1s ")

	q := &PagingQuery{
		Limit: limit,
	}
	p := GetModelPaginator(*q)
	fmt.Printf("Passou 2 ")

	result := p.Paginate(db, &dbRecords)
	if result.Error != nil {
		log.Panic("Erro")
	}
	fmt.Printf("Passou 3 ")

	// cursor := p.GetNextCursor()

	// pagination.Paging(&pagination.Param{
	// 	DB:      db,
	// 	Page:    *page,
	// 	Limit:   *limit,
	// 	OrderBy: []string{"id desc"},
	// }, &dbRecords)

	count := len(dbRecords)
	fmt.Printf("Passou count ")

	pageInfos := &models.PageInfo{
		BeforeCursor: p.GetNextCursor().Before,
		NextCursor:   p.GetNextCursor().After,
	}

	fmt.Printf("Passou 4 ")
	response := &models.Orders{
		Limit:    limit,
		Page:     page,
		Count:    &count,
		List:     dbRecords,
		PageInfo: pageInfos,
	}

	fmt.Printf("Passou 5 ")
	return response, db.Error
}
