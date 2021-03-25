package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

// ./Go-db_url 1 -jaeger_url 2 -kafka_broker 3 -port 4 -sentry_url 5 -some_app_id 6 -some_app_key 7

func Config() {
	var decision string
	fmt.Println("Please, write how do you want to config app?(json - write 'j'/flags or default - write 'f')")
	fmt.Scan(&decision)

	if decision == string('j') {
		jsonConfigure()
	} else {
		flagsConfigure()
	}
}

func flagsConfigure() {
	var (
		Port         = flag.Int("port", 8080, "port")
		DB_url       = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
		Jaeger_url   = flag.String("jaeger_url", "http://jaeger:16686", "jaeger url")
		Sentry_url   = flag.String("sentry_url", "http://sentry:9000", "sentry url")
		Kafka_broker = flag.String("kafka_broker", "kafka:9092", "kafka broker")
		Some_app_id  = flag.String("some_app_id", "testid", "some app id")
		Some_app_key = flag.String("some_app_key", "testkey", "some app key")
	)

	flag.Parse()

	urlValidation(*DB_url)
	urlValidation(*Jaeger_url)
	urlValidation(*Sentry_url)

	fmt.Printf("%d \n %s \n %s \n %s \n %s \n %s \n %s", *Port, *DB_url, *Jaeger_url, *Sentry_url, *Kafka_broker, *Some_app_id, *Some_app_key)
}

func jsonConfigure() {

	var config struct {
		Port         int    `json:"port"`
		DB_url       string `json:"db_url"`
		Jaeger_url   string `json:"jaeger_url"`
		Sentry_url   string `json:"sentry_url"`
		Kafka_broker string `json:"kafka_broker"`
		Some_app_id  string `json:"some_app_id"`
		Some_app_key string `json:"some_app_key"`
	}

	flag.Parse()

	if flag.Arg(0) == "" {
		log.Fatal(
			"Неверно задано количество аргументов командной строки. Проверьте, что вы задали имя файла.",
		)
	}

	filename := strings.TrimSpace(flag.Arg(0))

	configFile, err := os.Open(filename)
	fmt.Println()
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	urlValidation(config.DB_url)
	urlValidation(config.Jaeger_url)
	urlValidation(config.Sentry_url)

	fmt.Printf("%d \n %s \n %s \n %s \n %s \n %s \n %s", config.Port, config.DB_url, config.Jaeger_url, config.Sentry_url, config.Kafka_broker, config.Some_app_id, config.Some_app_key)
}

//Validations
func urlValidation(u string) error {
	_, err := url.ParseRequestURI(u)
	if err == nil {
		return nil
	} else {
		return err
	}
}
