package app2

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/arsu4ka/go-monorepo/pkg/config"
	"github.com/arsu4ka/go-monorepo/pkg/tasks"
	"github.com/hibiken/asynq"
	tele "gopkg.in/telebot.v3"
)

type BotOptions struct {
	RedisConfig config.Redis
	BotToken    string
	ListenPort  int
}

type Bot struct {
	tele        *tele.Bot
	asynqClient *asynq.Client
	srv         *fiber.App
	options     BotOptions
}

func NewBot(options BotOptions) *Bot {
	pref := tele.Settings{
		Token:     options.BotToken,
		Poller:    nil,
		ParseMode: tele.ModeHTML,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		tele:    b,
		options: options,
		asynqClient: asynq.NewClient(
			asynq.RedisClientOpt{Addr: options.RedisConfig.GetAddress()},
		),
		srv: fiber.New(),
	}
}

func (b *Bot) startHandler() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		var msg string
		if ctx.Sender().Username == "" {
			msg = "<b>Hello!</b>"
		} else {
			msg = fmt.Sprintf("<b>Hello, %s!</b>", ctx.Sender().Username)
		}

		task, err := tasks.NewUserJoined(fmt.Sprint(ctx.Sender().ID))
		if err != nil {
			log.Printf("Error creating task: %s\n", err.Error())
		} else {
			info, err := b.asynqClient.Enqueue(task, asynq.ProcessIn(20*time.Second))
			if err != nil {
				log.Printf("Error enqueuing task: %s\n", err.Error())
			}

			log.Printf("Created new task: %s\n", info.ID)
		}

		return ctx.Reply(msg)
	}
}

func (b *Bot) webhookHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var update tele.Update
		if err := c.BodyParser(&update); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		go b.tele.ProcessUpdate(update)
		return c.SendStatus(fiber.StatusOK)
	}
}

func (b *Bot) setupWebhookHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := c.Queries()
		webhookUrl, ok := query["url"]
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You should provide webhook url in query"})
		}

		telegramUrl := fmt.Sprintf(
			"https://api.telegram.org/bot%s/setWebhook?url=%s",
			b.options.BotToken,
			webhookUrl,
		)
		response, err := http.Get(telegramUrl)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		defer safeCloseBody(response.Body)

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error reading telegram response body: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to read telegram response}"})
		}

		return c.Status(response.StatusCode).Send(body)
	}
}

func (b *Bot) Start() error {
	b.tele.Handle("/start", b.startHandler())

	b.srv.Post("/webhook", b.webhookHandler())
	b.srv.Get("/telegram/setup", b.setupWebhookHandler())

	log.Print("Starting the bot")
	return b.srv.Listen(fmt.Sprintf(":%s", fmt.Sprint(b.options.ListenPort)))
}
