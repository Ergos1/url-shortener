package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Ergos1/url-shortener.git/infrastructure/db/psql"
	urlshortener "github.com/Ergos1/url-shortener.git/internal/app/url_shortener"
	"github.com/Ergos1/url-shortener.git/internal/config"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	cfg := config.NewConfig()

	db := psql.NewDB(ctx)
	if err := db.Connect(ctx, cfg.Database.Uri()); err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	urlRepo := urlshortener.NewUrlShortenerPsqlRepository(db)
	urlService := urlshortener.NewUrlShortenerService(urlRepo)

	result, err := urlService.CreateOrGetShortUrl(ctx, "https://www.google.com")
	fmt.Println(result, err)

	result, err = urlService.GetOriginalUrl(ctx, "8ffdefb")
	fmt.Println(result, err)

	return nil
}
