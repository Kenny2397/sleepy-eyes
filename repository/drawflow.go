package repository

import (
	"context"

	"github.com/Kenny2397/visual-programming/models"
)

type DrawflowRepository interface {
	InsertDrawflow(ctx context.Context, drawflow *models.Drawflow) error
	GetDrawflowById(ctx context.Context, id int64) (*models.Drawflow, error)
}

var implementation DrawflowRepository

func SetRepository(repository DrawflowRepository) {
	implementation = repository
}

func InsertDrawflow(ctx context.Context, drawflow *models.Drawflow) error {
	return implementation.InsertDrawflow(ctx, drawflow)
}

func GetDrawflowById(ctx context.Context, id int64) (*models.Drawflow, error) {
	return implementation.GetDrawflowById(ctx, id)
}
