package core

import "context"

type CreateOrGetShortUrlRequest struct {
	ShortUrl string `json:"short_url"`
}

func (s *Service) CreateOrGetShortUrl(ctx context.Context, url CreateOrGetShortUrlRequest) (string, error) {
	return s.urlShortenerService.CreateOrGetShortUrl(ctx, url.ShortUrl)
}
