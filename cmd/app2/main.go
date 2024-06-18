package main

import (
	"log"
	"os"
	"strconv"

	"github.com/arsu4ka/go-monorepo/internal/app2"
	"github.com/arsu4ka/go-monorepo/pkg/config"
)

func main() {
	redisConfig, err := config.GetRedisFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("APP2_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	botOptions := app2.BotOptions{
		RedisConfig: redisConfig,
		BotToken:    os.Getenv("BOT_TOKEN"),
		ListenPort:  port,
	}

	bot := app2.NewBot(botOptions)
	log.Fatal(bot.Start())
}
