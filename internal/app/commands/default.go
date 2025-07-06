package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) DefaultBehavior(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "you wrote: "+inputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID если нам нужно отвечать на какое то сообщение

	c.bot.Send(msg)
}
