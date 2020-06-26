package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/znobrega/go-gql-server/internal/gql"
	"github.com/znobrega/go-gql-server/internal/orm"
)

type Resolver struct {
	ORM *orm.ORM
}

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() gql.QueryResolver       { return &queryResolver{r} }

func (r *Resolver) Order() gql.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
