package resolvers

import (
	"context"
	"fmt"

	"github.com/znobrega/go-gql-server/internal/gql/models"
	log "github.com/znobrega/go-gql-server/internal/logger"
)

func (r *queryResolver) Transacoes(ctx context.Context, id *string, name *string) (*models.TransacoesFaturas, error) {
	if id != nil {
		fmt.Printf("id: %s \n", *id)
		log.Info("id: %s \n", *id)
	}

	if name != nil {
		fmt.Printf("name: %s \n", *name)
		log.Info("name: %s \n", *name)
	}

	transacao := &models.TransacoesFaturas{
		ID:          2,
		DsTransacao: "xand",
	}

	return transacao, nil
}
