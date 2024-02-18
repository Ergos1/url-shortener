package urlshortener

import "context"

type UrlRepository interface {
	Create(ctx context.Context, url Url) (int64, error)
	GetByOriginalUrl(ctx context.Context, originalUrl string) (string, error)
}

type UrlShortenerService struct {
	repo UrlRepository
}

func NewUrlShortenerService(repo UrlRepository) *UrlShortenerService {
	return &UrlShortenerService{
		repo: repo,
	}
}
