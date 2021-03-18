package main

import (
	"flag"

	"github.com/MelBogdan/Go/config"
)

func main() {
	flag.Parse()

	config.UrlValidation(*config.DB_url)
	config.UrlValidation(*config.Jaeger_url)
	config.UrlValidation(*config.Sentry_url)
}
