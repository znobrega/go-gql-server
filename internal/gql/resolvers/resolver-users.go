package resolvers

import (
	"context"

	"github.com/znobrega/go-gql-server/internal/gql/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	id := "ec17af15-e354-440c-a09f-69715fc8b595"
	email := "your@email.com"
	userId := "UserID-1"

	records := []*models.User{
		{
			ID:     id,
			Email:  email,
			UserID: userId,
		},
	}
	return records, nil
}
