package handlers

import (
	"context"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
)

type (
	photoHandler struct {
	}
)

func (ph *photoHandler) Endpoint() string {
	return telebot.OnPhoto
}

func (ph *photoHandler) Handle(
	ctx context.Context, msg *telebot.Message,
) {
	iBot := ctx.Value(CtxBotKey)
	bot := iBot.(*telebot.Bot)
	bot.Send(msg.Sender, "Received!")
	bot.Download(&msg.Photo.File, pathName(msg))
	fmt.Println(msg)
}

func NewPhotoHandler() Handler {
	return new(photoHandler)
}
