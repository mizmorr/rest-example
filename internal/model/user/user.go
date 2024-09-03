package user

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
	tableName struct{}  `pg:"users" gorm:"primaryKey"`
	ID        uuid.UUID `pg:"id,notnull, pk"`
	Firstname string    `pg:"firstname,notnull"`
	Lastname  string    `pg:"lastname,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

func (user *User) To_PGUser() *PGUser {

	return &PGUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
	}
}

func (pgUser *PGUser) To_User() *User {
	return &User{
		ID:        pgUser.ID,
		Firstname: pgUser.Firstname,
		Lastname:  pgUser.Lastname,
		CreatedAt: pgUser.CreatedAt,
	}
}
