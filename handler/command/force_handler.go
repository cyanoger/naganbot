package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/taranovegor/naganbot/router"
)

type ForceHandler struct {
	router.InterfaceCommandHandler
	bot *tgbotapi.BotAPI
}

func NewForceHandler(bot *tgbotapi.BotAPI) ForceHandler {
	return ForceHandler{
		bot: bot,
	}
}

func (handler ForceHandler) Name() string {
	return "ngforce"
}

func (handler ForceHandler) Execute(update tgbotapi.Update) error {
	_, err := handler.bot.Send(tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: update.FromChat().ID,
			UserID: update.SentFrom().ID,
		},
	})

	return err
}
