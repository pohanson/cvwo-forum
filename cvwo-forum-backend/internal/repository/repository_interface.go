package repository

import (
	"context"
	"sync"

	"github.com/pohanson/cvwo-forum/internal/database"
	"github.com/pohanson/cvwo-forum/internal/model"
)

type Repository struct {
	User UserRepository
}
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
}

var pgRepo *Repository
var once sync.Once

func GetPgRepo() *Repository {
	if pgRepo == nil {
		once.Do(func() {
			pgRepo = &Repository{
				User: &PostgresUserRepo{
					database.GetDb(),
				},
			}
		})
	}
	return pgRepo
}
