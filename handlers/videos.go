package handlers

import (
	"context"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
)

type (
	videosHandler struct {
	}
)

func (vh *videosHandler) Endpoint() string{
	return telebot.OnVideo
}

func (vh *videosHandler) Handle(ctx context.Context, msg *telebot.Message){
	iBot := ctx.Value(CtxBotKey)
	bot := iBot.(*telebot.Bot)
	bot.Send(msg.Sender, "Received!")
	bot.Download(&msg.Video.File, pathName(msg))
	fmt.Println(msg)
}

func NewVideoHandler() Handler {
	return new(videosHandler)
}
