package urlshortener

import (
	"context"
	"errors"

	"github.com/Ergos1/url-shortener.git/infrastructure/db/psql"
	"github.com/jackc/pgx/v4"
)

type UrlShortenerPsqlRepository struct {
	db psql.PGX
}

func NewUrlShortenerPsqlRepository(db psql.PGX) *UrlShortenerPsqlRepository {
	return &UrlShortenerPsqlRepository{
		db: db,
	}
}

func (r *UrlShortenerPsqlRepository) Create(ctx context.Context, url Url) (int64, error) {
	var id int64
	err := r.db.Create(ctx, &id, "INSERT INTO urls(short_url, original_url) VALUES($1, $2) RETURNING id", url.ShortUrl, url.OriginalUrl)
	return id, err
}

func (r *UrlShortenerPsqlRepository) GetByOriginalUrl(ctx context.Context, originalUrl string) (*UrlRow, error) {
	var url UrlRow
	err := r.db.Get(ctx, &url, "SELECT * FROM urls WHERE original_url = $1", originalUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUrlNotFound
		}

		return nil, err
	}

	return &url, err
}

func (r *UrlShortenerPsqlRepository) GetByShortUrl(ctx context.Context, shortUrl string) (*UrlRow, error) {
	var url UrlRow
	err := r.db.Get(ctx, &url, "SELECT * FROM urls WHERE short_url = $1", shortUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUrlNotFound
		}

		return nil, err
	}

	return &url, err
}
