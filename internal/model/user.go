package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
type PGUser struct {
	ID        uuid.UUID `pg:"id,notnull, pk"`
	Firstname string    `pg:"firstname,notnull"`
	Lastname  string    `pg:"lastname,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

type UserCreateRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

type UserUpdateRequest struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
}

func (user *User) ToPg() *PGUser {

	return &PGUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
	}
}

func (pgUser *PGUser) ToWeb() *User {
	return &User{
		ID:        pgUser.ID,
		Firstname: pgUser.Firstname,
		Lastname:  pgUser.Lastname,
		CreatedAt: pgUser.CreatedAt,
	}
}

func (ucr *UserCreateRequest) ToPg() *PGUser {
	return &PGUser{
		Firstname: ucr.Firstname,
		Lastname:  ucr.Lastname,
	}
}

func (upr *UserUpdateRequest) ToPg() *PGUser {
	return &PGUser{
		ID:        upr.ID,
		Firstname: upr.Firstname,
		Lastname:  upr.Lastname,
	}
}
