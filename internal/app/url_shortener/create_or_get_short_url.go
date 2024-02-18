package urlshortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

func generateShortUrl(url string) (string, error) {
	md5Hash := md5.New()
	if _, err := md5Hash.Write([]byte(url)); err != nil {
		return "", err
	}

	hashResult := hex.EncodeToString(md5Hash.Sum(nil))
	return hashResult[:7], nil
}

func (s *UrlShortenerService) CreateOrGetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	urlRow, err := s.repo.GetByOriginalUrl(ctx, originalUrl)
	if err != nil && err != ErrUrlNotFound {
		return "", err
	}

	if urlRow != nil {
		return urlRow.ShortUrl, nil
	}

	// TODO: use bloom filter to check if the short url is already in use
	var shortUrl string
	urlToHash := originalUrl
	for {
		shortUrl, err = generateShortUrl(urlToHash)
		if err != nil {
			return "", err
		}

		urlRow, err = s.repo.GetByShortUrl(ctx, shortUrl)
		if err != nil && err != ErrUrlNotFound {
			return "", err
		}

		if urlRow == nil {
			break
		}

		urlToHash += HASH_PREDEFINED_STRING
	}

	url := Url{
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
	}

	if _, err = s.repo.Create(ctx, url); err != nil {
		return "", err
	}

	return shortUrl, nil

}
