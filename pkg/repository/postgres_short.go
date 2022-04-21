package repository

import (
	"api"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ShortRepository struct {
	db *sqlx.DB
}

func NewShortRepository(db *sqlx.DB) *ShortRepository {
	return &ShortRepository{
		db: db,
	}
}
func (r *ShortRepository) Create(user api.Short) (string, error) {

	var shorturl string
	qwery := fmt.Sprintf("insert into %s(short_code, long_url, short_url) values ($1, $2, $3) returning short_url", ShortTable)
	row := r.db.QueryRow(qwery, user.ShortCode, user.LongURL, user.ShortURL)
	if err := row.Scan(&shorturl); err != nil {
		return "", err
	}
	return shorturl, nil
}

func (r *ShortRepository) GetUrl(shortUrl string) (string, error) {
	var longurl string
	qwery := fmt.Sprintf("select long_url from %s where short_code=($1)", ShortTable)
	if err := r.db.Get(&longurl, qwery, shortUrl); err != nil {
		return longurl, err
	}
	return longurl, nil
}
