package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/chigvero/TelegramBot/internal/app/commands"
	"github.com/chigvero/TelegramBot/internal/service/product"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productServce := product.NewService()

	commander := commands.NewCommander(bot, productServce)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.HelpCommand(update.Message)
		case "list":
			commander.ListCommand(update.Message)
		default:
			commander.DefaultBehavior(update.Message)
		}
	}
}
