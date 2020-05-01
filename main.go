package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)

		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, fmt.Sprintf(hello, m.Sender.FirstName), tb.ModeHTML); err != nil {

		}
	})

	b.Handle("/help", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, help, tb.ModeHTML); err != nil {

		}
	})

	b.Handle("/search", func(m *tb.Message) {
		if _, err := b.Send(m.Sender, getReply(m.Text), tb.ModeHTML); err != nil {
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
