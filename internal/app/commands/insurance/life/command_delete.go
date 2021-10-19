package life

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const deleteSignature = "method signature: /delete__insurance__life {uint}"
const deleteErrorMassage = "incorrect input\n\n" + deleteSignature

func (telegramLifeCommander *TelegramLifeCommander) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil || idx < 0 {
		telegramLifeCommander.sendError(inputMessage, deleteErrorMassage)
		return
	}

	if result, err := telegramLifeCommander.lifeService.Remove(uint64(idx)); result {
		telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Success"))
	} else {
		telegramLifeCommander.sendError(inputMessage, err.Error())
	}

}
