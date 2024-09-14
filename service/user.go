package service

import (
	"context"
	"fmt"

	"github.com/mizmorr/rest-example/internal/model"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/store"
)

type UserWebService struct {
	store *store.Store
	ctx   context.Context
}

func NewUserWebService(store *store.Store, ctx context.Context) (*UserWebService, error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}
	return &UserWebService{
		store: store,
		ctx:   ctx,
	}, nil
}

func (svc *UserWebService) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {

	user, err := svc.store.User.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Get")
	}
	if user == nil {
		return nil, errors.New(fmt.Sprintf("User %v not found", id))
	}
	return user.ToWeb(), nil
}

func (svc *UserWebService) CreateUser(ctx context.Context, reqUser *model.UserCreateRequest) (*model.User, error) {

	id, err := svc.store.User.Create(ctx, reqUser.ToPg())

	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Create")
	}

	createdUser, err := svc.store.User.Get(ctx, id)

	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Create")
	}

	return createdUser.ToWeb(), nil
}

func (svc *UserWebService) DeleteUser(ctx context.Context, id uuid.UUID) error {

	err := svc.store.User.Delete(ctx, id)

	if err != nil {
		return errors.Wrap(err, "svc.User.Delete")
	}

	return nil
}

func (svc *UserWebService) UpdateUser(ctx context.Context, req *model.UserUpdateRequest) (*model.User, error) {

	id, err := svc.store.User.Update(ctx, req.ToPg())

	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Update")
	}

	updatedUser, err := svc.store.User.Get(ctx, id)

	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Update")
	}

	if updatedUser.Firstname != req.Firstname || updatedUser.Lastname != req.Lastname {
		return nil, errors.Wrap(fmt.Errorf("user not updated"), "svc.User.Update")
	}

	return updatedUser.ToWeb(), nil
}
