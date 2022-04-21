package repository

import (
	"api"
	"github.com/jmoiron/sqlx"
)

type Short interface {
	Create(url api.Short) (string, error)
	GetUrl(shortUrl string) (string, error)
}

type Repository struct {
	Short
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Short: NewShortRepository(db),
	}
}
