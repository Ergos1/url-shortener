package core

import "context"

func (s *Service) GetOriginalUrl(ctx context.Context, shortURL string) (string, error) {
	return s.urlShortenerService.GetOriginalUrl(ctx, shortURL)
}
