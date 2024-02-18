package core

type UrlShortenerService interface {
	CreateOrGetShortUrl(url string) (string, error)
	GetOriginalUrl(shortUrl string) (string, error)
}

type Service struct {
	urlShortenerService UrlShortenerService
}

func NewService(urlShortenerService UrlShortenerService) *Service {
	return &Service{
		urlShortenerService: urlShortenerService,
	}
}
