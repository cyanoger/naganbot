package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type CommandRouter struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]InterfaceCommandHandler
}

type InterfaceCommandHandler interface {
	Name() string
	Execute(update tgbotapi.Update) error
}

func NewCommandRouter(bot *tgbotapi.BotAPI, handlers ...InterfaceCommandHandler) CommandRouter {
	router := CommandRouter{
		bot:      bot,
		handlers: map[string]InterfaceCommandHandler{},
	}

	for _, handler := range handlers {
		router.attachHandler(handler)
	}

	return router
}

func (router CommandRouter) attachHandler(handler InterfaceCommandHandler) {
	router.handlers[handler.Name()] = handler
}

func (router CommandRouter) Navigate(update tgbotapi.Update) {
	if nil == update.Message {
		return
	}

	if !update.Message.IsCommand() {
		return
	}

	err := router.handlers[update.Message.Command()].Execute(update)
	if err != nil {
		log.Printf("Router: returned error during command `%s` execution: %s", update.Message.Command(), err.Error())
	}
}
