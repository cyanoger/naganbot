package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/taranovegor/naganbot/handler/command"
	"github.com/taranovegor/naganbot/router"
	"log"
	"os"
)

func main() {
	fmt.Println("Naganbot / Shoot yourself in Russian roulette")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("App: error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("APP_TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("App: authorized on account %s", bot.Self.UserName)

	if "true" == os.Getenv("APP_DEBUG") {
		bot.Debug = true
	} else {
		bot.Debug = false
	}

	r := router.NewCommandRouter(
		bot,
		command.NewForceHandler(bot),
	)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updateConfig.AllowedUpdates = []string{"message"}

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		go r.Navigate(update)
	}
}
