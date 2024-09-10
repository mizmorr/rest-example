package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model"
	"github.com/mizmorr/rest-example/store/pg"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Get(ctx context.Context, id uuid.UUID) (*model.PGUser, error) {

	user := model.PGUser{}
	query := `
		select * from users where id=$1
	`
	row := repo.db.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Create(ctx context.Context, user *model.PGUser) error {
	var id uuid.UUID
	query := `
	input into users values($1,$2,$3) returning id
	`
	err := repo.db.QueryRow(ctx, query, user.ID, user.Firstname, user.Lastname).Scan(&id)

	if err != nil {
		return err
	}
	return nil
}
