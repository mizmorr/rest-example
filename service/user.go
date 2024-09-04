package service

import (
	"context"
	"fmt"

	"github.com/mizmorr/rest-example/internal/model/user"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/store"
)

type UserWebService struct {
	store *store.Store
	ctx   context.Context
}

func NewUserWebService(store *store.Store, ctx context.Context)  (*UserWebService,error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}
	return &UserWebService{
		store: store,
		ctx:   ctx,
	}, nil
}

func (svc *UserWebService) GetUser(ctx context.Context, id uuid.UUID) (*user.User, error) {

	user, err := svc.store.User.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Get")
	}
	if user == nil {
		return nil, errors.New(fmt.Sprintf("User %v not found", id))
	}
	return user.ToWeb(), nil
}
