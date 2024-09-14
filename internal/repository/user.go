package repository

import (
	"context"
	"fmt"

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

func (repo *UserRepo) Create(ctx context.Context, user *model.PGUser) (uuid.UUID, error) {
	var (
		newUserID  uuid.UUID = uuid.New()
		returnedID uuid.UUID
	)

	query := `
	insert into users values($1,$2,$3) returning id
	`

	err := repo.db.QueryRow(ctx, query, newUserID, user.Firstname, user.Lastname).Scan(&returnedID)

	if err != nil {
		return uuid.UUID{}, err
	}
	return returnedID, nil
}

func (repo *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {

	query := `
	delete from users where id=$1
	`
	res, err := repo.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (repo *UserRepo) Update(ctx context.Context, user *model.PGUser) (uuid.UUID, error) {

	var id uuid.UUID

	query := `
	update users set firstname=$1, lastname=$2 where id=$3 returning id
	`
	row := repo.db.QueryRow(ctx, query, user.Firstname, user.Lastname, user.ID)

	if err := row.Scan(&id); err != nil {
		return uuid.UUID{}, err

	}

	return id, nil
}

func (repo *UserRepo) GetAll(ctx context.Context) ([]*model.PGUser, error) {

	rows, err := repo.db.Query(ctx, "select * from users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.PGUser
	// to do allocation users := make([]model.PGUser, 0, selected_users)

	for rows.Next() {
		var user model.PGUser
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
