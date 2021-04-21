package handlers

import (
	"context"
	"os"
	"strconv"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

const (
	CtxBotKey = "bot"
)

type (
	HandlerFunc func(ctx context.Context, msg *telebot.Message)
)

func Adapt(handlerFunc HandlerFunc, bot *telebot.Bot) func(message *telebot.Message) {
	return func(message *telebot.Message) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, CtxBotKey, bot)
		handlerFunc(ctx, message)
	}
}

func pathName(msg *telebot.Message) string {
	senderID := strconv.Itoa(msg.Sender.ID)
	dirPath := os.Getenv("STORAGE_PATH") + "/" + senderID
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if nil != err {
			panic("error creating directory!!")
		}
	}
	time := time.Now()
	layout := "2006-01-02"
	date := time.Format(layout)

	switch true {
	case msg.Document != nil:
		return dirPath + "/" + msg.Document.FileName
	case msg.Photo != nil:
		return dirPath + "/" + msg.Photo.UniqueID
	case msg.Video != nil:
		return dirPath + "/" + msg.Video.FileName
	case msg.Text != "":
		return dirPath + "/" + date + ".txt"
	default:
		panic("unmapped case")
	}
}
