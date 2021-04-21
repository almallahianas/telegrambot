package server

import (
	"github.com/almallahianas/telegrambot/handlers"
	"gopkg.in/tucnak/telebot.v2"
)

type (
	Server interface {
		Start()
		Stop()
		RegisterHandler(h handlers.Handler)
	}

	server struct {
		bot *telebot.Bot
	}
)

func (s *server) Start() {
	s.bot.Start()
}

func (s *server) Stop() {
	s.bot.Stop()
}

func (s *server) RegisterHandler(h handlers.Handler) {
	s.bot.Handle(h.Endpoint(), handlers.Adapt(h.Handle, s.bot))
}

func NewServer(bot *telebot.Bot) Server {
	return &server{bot: bot}
}
