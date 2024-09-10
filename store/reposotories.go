package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model"
)

type UserRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*model.PGUser, error)
	Create(ctx context.Context, user *model.PGUser) error
}
