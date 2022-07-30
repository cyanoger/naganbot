package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

type testHandler struct {
	InterfaceCommandHandler
}

func (handler testHandler) Name() string {
	return "test"
}

func (handler testHandler) Execute(update tgbotapi.Update) error {
	return nil
}

func TestNewCommandRouter(t *testing.T) {
	handler := testHandler{}
	router := NewCommandRouter(new(tgbotapi.BotAPI), handler)

	if routerHandler, handlerExists := router.handlers[handler.Name()]; !handlerExists || handler.Name() != routerHandler.Name() {
		t.Fail()
	}
}
