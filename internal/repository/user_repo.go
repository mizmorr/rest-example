package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model/user"
	"github.com/mizmorr/rest-example/store/pg"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Get_User(ctx context.Context, id uuid.UUID) (*user.PGUser, error) {

	user := user.PGUser{}
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
