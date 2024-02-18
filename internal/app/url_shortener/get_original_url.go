package urlshortener

import (
	"context"
)

func (s *UrlShortenerService) GetOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.repo.GetByShortUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	if url == nil {
		return "", nil
	}

	return url.OriginalUrl, nil
}
