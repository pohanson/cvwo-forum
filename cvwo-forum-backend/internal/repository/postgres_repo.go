package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/pohanson/cvwo-forum/internal/model"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func (repo *PostgresUserRepo) Create(ctx context.Context, user *model.User) (model.User, error) {
	u := model.User{}
	err := repo.db.QueryRowContext(ctx, "INSERT INTO appuser(username, name, role) VALUES($1, $2, $3) RETURNING id, username, name, role", user.Username, user.Name, user.Role).Scan(&u.Id, &u.Username, &u.Name, &u.Role)
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
	err := repo.db.QueryRowContext(ctx, "SELECT id, username, name, role FROM appuser WHERE username=$1;", username).Scan(&u.Id, &u.Username, &u.Name, &u.Role)
	return u, err
}

type PostgresThreadRepo struct {
	db *sql.DB
}

func (repo *PostgresThreadRepo) Create(ctx context.Context, thread *model.Thread) (*model.Thread, error) {
	t := model.Thread{}
	err := repo.db.QueryRowContext(ctx, "INSERT INTO thread(title, content, type, created_by) VALUES($1, $2, $3, $4) RETURNING id, title, content, type, created_by, created_on, edited_on;", thread.Title, thread.Content, thread.Type, thread.CreatedBy).Scan(&t.Id, &t.Title, &t.Content, &t.Type, &t.CreatedBy, &t.CreatedOn, &t.EditedOn)
	return &t, err
}

type PostgresCategoryRepo struct {
	db *sql.DB
}

func (repo *PostgresCategoryRepo) Create(ctx context.Context, category *model.Category) (*model.Category, error) {
	c := model.Category{}
	err := repo.db.QueryRowContext(ctx, "INSERT INTO category(title, description) VALUES($1, $2) RETURNING id, title, description;", category.Title, category.Description).Scan(&c.Id, &c.Title, &c.Description)
	return &c, err
}

func (repo *PostgresCategoryRepo) GetAll(ctx context.Context) (*[]model.Category, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, description FROM category;")
	if err != nil {
		return nil, err
	}
	results := []model.Category{}
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.Id, &category.Title, &category.Description)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, category)
	}
	return &results, nil
}

type PostgresThreadCategoryRepo struct {
	db                 *sql.DB
	threadRepo         ThreadRepository
	categoryRepository CategoryRepository
}

func (repo *PostgresThreadCategoryRepo) Create(ctx context.Context, thread *model.Thread, category *model.Category) error {
	err := repo.db.QueryRowContext(ctx, "INSERT INTO thread_category(thread_id, category_id) VALUES($1, $2);", thread.Id, category.Id).Err()
	return err
}
