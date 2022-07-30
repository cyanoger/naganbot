package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func newForceHandler() ForceHandler {
	return NewForceHandler(new(tgbotapi.BotAPI))
}

func TestForceHandler(t *testing.T) {
	handler := newForceHandler()

	if nil == handler.bot {
		t.Fail()
	}
}

func TestName(t *testing.T) {
	if "ngforce" != newForceHandler().Name() {
		t.Fail()
	}
}
