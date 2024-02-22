package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Ergos1/url-shortener.git/infrastructure/db/psql"
	"github.com/Ergos1/url-shortener.git/internal/app/core"
	urlshortener "github.com/Ergos1/url-shortener.git/internal/app/url_shortener"
	"github.com/Ergos1/url-shortener.git/internal/config"
	"github.com/Ergos1/url-shortener.git/internal/controllers/http"
	"github.com/Ergos1/url-shortener.git/internal/controllers/http/handlers"
)

type Server interface {
	Run() error
}

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
	coreService := core.NewService(urlService)

	var server Server = http.NewServer(ctx,
		http.WithAddress(cfg.Server.Address),
		http.WithMount("/", handlers.NewBaseHandler().Routes()),
		http.WithMount("/url", handlers.NewUrlShortenerHandler(coreService).Routes()),
	)

	if err := server.Run(); err != nil {
		log.Printf("[MAIN] Error while running server: %v", err)
	}

	return nil
}
