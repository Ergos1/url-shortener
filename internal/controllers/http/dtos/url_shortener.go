package dtos

import "github.com/Ergos1/url-shortener.git/internal/app/core"

type CreateOrGetShortUrlRequest struct {
	Url string `json:"url" validate:"required"`
}

type CreateOrGetShortUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type GetOriginalUrlResponse struct {
	OriginalUrl string `json:"originalUrl"`
}

func (r CreateOrGetShortUrlRequest) ToServiceRequest() core.CreateOrGetShortUrlRequest {
	return core.CreateOrGetShortUrlRequest{
		ShortUrl: r.Url,
	}
}
