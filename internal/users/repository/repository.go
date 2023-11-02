package repository

import (
	"context"
	"database/sql"

	"github.com/krlosaren/hostal-app/internal/users/entities"
)

type (
	Repository interface {
		SaveNewUser(context.Context, entities.User) (int64, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(c *sql.DB) Repository {
	return &repository{
		db: c,
	}
}

func (repository *repository) SaveNewUser(ctx context.Context, user entities.User) (int64, error) {
}
