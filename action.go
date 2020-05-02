package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var (
	publicURL = os.Getenv("PUBLIC_URL")
	port      = os.Getenv("PORT")
	token     = os.Getenv("TOKEN")
)

type BotBuilder interface {
	getBot() *tb.Bot
}

type Poller struct{}

type WebHook struct{}

func (p Poller) getBot() *tb.Bot {

	log.Print("Creating poller bot...")
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}

	return b
}

func (w WebHook) getBot() *tb.Bot {

	log.Print("Creating Webhook bot...")
	b, err := tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.Webhook{
			Listen:   ":" + port,
			Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
		},
	})

	if err != nil {
		panic(err)
	}

	return b
}
