package core

import "context"

type UrlShortenerService interface {
	CreateOrGetShortUrl(ctx context.Context, url string) (string, error)
	GetOriginalUrl(ctx context.Context, shortUrl string) (string, error)
}

type Service struct {
	urlShortenerService UrlShortenerService
}

func NewService(urlShortenerService UrlShortenerService) *Service {
	return &Service{
		urlShortenerService: urlShortenerService,
	}
}
