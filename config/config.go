package config

import (
	"flag"
	"log"
	"net/url"
)

var (
	Port         = flag.Int("port", 8080, "port")
	DB_url       = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
	Jaeger_url   = flag.String("jaeger_url", "http://jaeger:16686", "jaeger url")
	Sentry_url   = flag.String("sentry_url", "http://sentry:9000", "sentry url")
	Kafka_broker = flag.String("kafka_broker", "kafka:9092", "kafka broker")
	Some_app_id  = flag.String("some_app_id", "testid", "some app id")
	Some_app_key = flag.String("some_app_key", "testkey", "some app key")
)

// ./main -db_url 1 -jaeger_url 2 -kafka_broker 3 -port 4 -sentry_url 5 -some_app_id 6 -some_app_key 7

func UrlValidation(u string) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		log.Fatal("Некорректно указан url")
	}
}
