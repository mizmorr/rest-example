package user

type PGUser struct {
	tableName struct{} `pg:"users"`
	ID        int64    `pg:"id,notnull, pk"`
	Firstname string   `pg:"firstname,notnull"`
	Lastname  string   `pg:"lastname,notnull"`
	CreatedAt string   `pg:"created_at,notnull"`
}

func (user *User) To_PGUser() *PGUser {

	return &PGUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
	}
}
