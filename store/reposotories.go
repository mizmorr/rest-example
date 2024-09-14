package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model"
)

type UserRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*model.PGUser, error)
	Create(ctx context.Context, user *model.PGUser) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, user *model.PGUser) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]*model.PGUser, error)
}
