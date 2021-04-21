package main

import (
	"os"
	"time"

	"github.com/almallahianas/telegrambot/handlers"
	server2 "github.com/almallahianas/telegrambot/server"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	botSettings := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(botSettings)
	if nil != err {
		panic(err)
	}

	server := server2.NewServer(bot)
	server.RegisterHandler(handlers.NewTextHandler())
	server.RegisterHandler(handlers.NewDocumentHandler())
	server.RegisterHandler(handlers.NewPhotoHandler())
	server.RegisterHandler(handlers.NewVideoHandler())
	server.Start()
}
