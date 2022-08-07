package router

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type CommandRouter struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]InterfaceCommandHandler
}

type InterfaceCommandHandler interface {
	Name() string
	Execute(tgbotapi.Update) error
}

func NewCommandRouter(bot *tgbotapi.BotAPI, handlers ...InterfaceCommandHandler) CommandRouter {
	router := CommandRouter{
		bot:      bot,
		handlers: map[string]InterfaceCommandHandler{},
	}

	for _, handler := range handlers {
		router.addHandler(handler)
	}

	return router
}

func (router CommandRouter) addHandler(handler InterfaceCommandHandler) {
	router.handlers[handler.Name()] = handler
}

func (router CommandRouter) getHandler(command string) (InterfaceCommandHandler, error) {
	handler, exist := router.handlers[command]
	if exist {
		return nil, errors.New(fmt.Sprintf("Handler for `%s` command not found", command))
	}

	return handler, nil
}

func (router CommandRouter) Navigate(update tgbotapi.Update) {
	if nil == update.Message || !update.Message.IsCommand() {
		return
	}

	handler, err := router.getHandler(update.Message.Command())
	if err != nil {
		return
	}

	err = handler.Execute(update)
	if err != nil {
		log.Printf("Router: Returned error during command `%s` execution: %s", update.Message.Command(), err.Error())
	}
}
