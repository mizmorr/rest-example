package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model/user"
)

type UserRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*user.PGUser, error)
}
