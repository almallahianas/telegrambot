package handlers

import (
	"context"
	"gopkg.in/tucnak/telebot.v2"
	"os"
)

type (
	Handler interface {
		Endpoint() string
		Handle(ctx context.Context, msg *telebot.Message)
	}

	textHandler struct {
	}
)

func (t *textHandler) Endpoint() string {
	return telebot.OnText
}

func (t *textHandler) Handle(
	ctx context.Context, msg *telebot.Message,
) {
	iBot := ctx.Value(CtxBotKey)
	bot := iBot.(*telebot.Bot)
	bot.Send(msg.Sender, "Saved Your Chat ")
	filePath := pathName(msg)

	var file, err = os.OpenFile(
		filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)

	if nil != err {
		panic("Cant open the file to write")
	} else {
		_, err1 := file.WriteString(msg.Text + "\n")
		if nil != err1 {
			panic("cant write on the file")
		}
		err2 := file.Sync()
		if err2 != nil {
			panic("changes on the file are not saved")
		}
	}
}

func NewTextHandler() Handler {
	return new(textHandler)
}
