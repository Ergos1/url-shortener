package urlshortener

type Url struct {
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}

func (u *Url) MapFromModel(url *UrlRow) *Url {
	return &Url{
		ShortUrl:    url.ShortUrl,
		OriginalUrl: url.OriginalUrl,
	}
}

func (u *Url) MapToModel(id int64) *UrlRow {
	return &UrlRow{
		ID:          id,
		ShortUrl:    u.ShortUrl,
		OriginalUrl: u.OriginalUrl,
	}
}
