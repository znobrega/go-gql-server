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

	// p.SetKeys("customer_name")

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

// Paging workin with cursor after page
// paginator "github.com/pilagod/gorm-cursor-paginator"
func (r *queryResolver) Orders(ctx context.Context, first *int, page *int, filter map[string]interface{}, after *string, before *string) (*models.Orders, error) {
	var dbRecords []*models.Order
	var cursors []string
	var edges []*models.EdgeOrder

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
	sort := "asc"
	q := &PagingQuery{
		After:  after,
		Before: before,
		Limit:  first,
		Order:  &sort,
	}
	p := GetModelPaginator(*q)
	fmt.Printf("Passou 2 ")

	result := p.Paginate(db, &dbRecords)
	if result.Error != nil {
		log.Panic("Erro")
	}
	fmt.Printf("Passou 3 ")

	count := len(dbRecords)
	fmt.Printf("Passou count ")

	pageInfos := &models.PageInfo{
		BeforeCursor:  nil,
		NextCursor:    nil,
		StartCursor:   p.GetNextCursor().Before,
		EndCursor:     p.GetNextCursor().After,
		HasNextPage:   models.HasPage(p.GetNextCursor().After),
		HasBeforePage: models.HasPage(p.GetNextCursor().Before),
	}

	cursors = p.GetCursors()

	for i, element := range cursors {
		edge := &models.EdgeOrder{
			Node:   *dbRecords[i],
			Cursor: element,
		}

		edges = append(edges, edge)
	}

	fmt.Printf("Passou 4 ")
	response := &models.Orders{
		Limit:    first,
		Page:     page,
		Count:    &count,
		Edges:    edges,
		PageInfo: pageInfos,
	}

	fmt.Printf("Passou 5 ")
	return response, db.Error
}
