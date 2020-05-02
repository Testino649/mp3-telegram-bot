package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

func main() {

	var bb BotBuilder

	if os.Getenv("environment") == "dev" {
		bb = Poller{}
	} else {
		bb = WebHook{}
	}

	b := getBot(bb)

	b.Handle("/start", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, fmt.Sprintf(welcome, m.Sender.FirstName), tb.ModeHTML); err != nil {

		}
	})

	b.Handle("/hello", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, fmt.Sprintf(hello, m.Sender.FirstName), tb.ModeHTML); err != nil {

		}
	})

	b.Handle("/help", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, help, tb.ModeHTML); err != nil {

		}
	})

	b.Handle("/search", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, getReply(m.Payload), tb.ModeHTML); err != nil {
			log.Println(err)
		}
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if _, err := b.Send(m.Sender, allReply, tb.ModeHTML); err != nil {
			log.Println(err)
		}
	})

	b.Start()

	log.Println("bot running....")

}

func getBot(bb BotBuilder) *tb.Bot {

	return bb.getBot()

}
