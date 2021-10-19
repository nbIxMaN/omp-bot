package life

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/insurance"
	"strconv"
)

const newSignature = "method signature: /new__insurance__life {life in JSON}"
const newErrorMassage = "incorrect input\n\n" + newSignature

func (telegramLifeCommander *TelegramLifeCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	var life insurance.Life
	err := json.Unmarshal([]byte(args), &life)
	if err != nil {
		telegramLifeCommander.sendError(inputMessage, newErrorMassage)
		return
	}

	if result, err := telegramLifeCommander.lifeService.Create(life); err == nil {
		telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID,
			"ID: "+strconv.FormatUint(result, 10)))
	} else {
		telegramLifeCommander.sendError(inputMessage, err.Error())
	}
}
