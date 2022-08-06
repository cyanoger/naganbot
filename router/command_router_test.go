package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

const (
	TestHandlerName = "test"
)

type testHandler struct {
	InterfaceCommandHandler
}

func (handler testHandler) Name() string {
	return TestHandlerName
}

func (handler testHandler) Execute(update tgbotapi.Update) error {
	return nil
}

func newTestCommandRouter() CommandRouter {
	handler := testHandler{}
	router := NewCommandRouter(new(tgbotapi.BotAPI), handler)

	return router
}

func TestNewCommandRouter(t *testing.T) {
	router := newTestCommandRouter()

	if routerHandler, handlerExists := router.handlers[TestHandlerName]; !handlerExists || TestHandlerName != routerHandler.Name() {
		t.Fail()
	}
}

func TestGetHandler(t *testing.T) {
	router := newTestCommandRouter()

	handler, err := router.getHandler(TestHandlerName)
	if err != nil {
		t.Fail()
	}

	if TestHandlerName != handler.Name() {
		t.Fail()
	}
}
