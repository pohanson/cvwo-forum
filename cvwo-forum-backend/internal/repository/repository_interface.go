package repository

import (
	"context"
	"sync"

	"github.com/pohanson/cvwo-forum/internal/database"
	"github.com/pohanson/cvwo-forum/internal/model"
)

type Repository struct {
	User           UserRepository
	Thread         ThreadRepository
	Category       CategoryRepository
	ThreadCategory ThreadCategoryRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
}

type ThreadRepository interface {
	Create(ctx context.Context, thread *model.Thread) (*model.Thread, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category *model.Category) (*model.Category, error)
	GetAll(ctx context.Context) (*[]model.Category, error)
}

type ThreadCategoryRepository interface {
	Create(ctx context.Context, thread *model.Thread, category *model.Category) error
}

var pgRepo *Repository
var once sync.Once

func GetPgRepo() *Repository {
	if pgRepo == nil {
		once.Do(func() {
			db := database.GetDb()

			userRepo := &PostgresUserRepo{db}
			threadRepo := &PostgresThreadRepo{db}
			categoryRepo := &PostgresCategoryRepo{db}
			threadCategoryRepo := &PostgresThreadCategoryRepo{db, threadRepo, categoryRepo}

			pgRepo = &Repository{
				User:           userRepo,
				Thread:         threadRepo,
				Category:       categoryRepo,
				ThreadCategory: threadCategoryRepo,
			}
		})
	}
	return pgRepo
}
