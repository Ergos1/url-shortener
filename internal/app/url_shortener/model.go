package urlshortener

type UrlRow struct {
	ID          int64  `db:"id"`
	ShortUrl    string `db:"short_url"`
	OriginalUrl string `db:"original_url"`
}
