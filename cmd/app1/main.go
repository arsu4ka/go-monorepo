package main

import (
	"log"

	"github.com/arsu4ka/go-monorepo/internal/app1"
	"github.com/arsu4ka/go-monorepo/pkg/config"
)

func main() {
	redisConfig, err := config.GetRedisFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := app1.NewServer(redisConfig)
	log.Fatal(srv.Start())
}
