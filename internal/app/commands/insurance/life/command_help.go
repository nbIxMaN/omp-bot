package life

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (telegramLifeCommander *TelegramLifeCommander) Help(inputMessage *tgbotapi.Message) {

	responseMessage := "help. /help__insurance__life\n" +
		"list life. /list__insurance__life\n" +
		"delete life. " + deleteSignature + "\n" +
		"edit life. " + editSignature + "\n" +
		"get life. " + getSignature + "\n" +
		"new life. " + newSignature

	telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, responseMessage))

}
