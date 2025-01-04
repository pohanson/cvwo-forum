package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/pohanson/cvwo-forum/internal/model"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func (repo *PostgresUserRepo) Create(ctx context.Context, user *model.User) (model.User, error) {
	u := model.User{}
	err := repo.db.QueryRowContext(ctx, "INSERT INTO appuser(username, name, role) VALUES($1, $2, $3) RETURNING username, name, role", user.Username, user.Name, user.Role).Scan(&u.Username, &u.Name, &u.Role)
	if pqerr, ok := err.(*pq.Error); ok {
		switch pqerr.Code {
		case "23505":
			return u, &UniqueViolationErr{"Username"}
		}
	}
	return u, err
}

func (repo *PostgresUserRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	u := model.User{}
	err := repo.db.QueryRowContext(ctx, "SELECT username, name, role FROM appuser WHERE username=$1", username).Scan(&u.Username, &u.Name, &u.Role)
	return u, err
}
