package life

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/insurance"
)

type LifeCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type LifeService interface {
	Describe(LifeID uint64) (*insurance.Life, error)
	List(cursor uint64, limit uint64) ([]insurance.Life, error)
	Create(life insurance.Life) (uint64, error)
	Update(LifeID uint64, life insurance.Life) error
	Remove(LifeID uint64) (bool, error)
}

type TelegramLifeCommander struct {
	bot         *tgbotapi.BotAPI
	lifeService LifeService
}

func NewTelegramLifeCommander(bot *tgbotapi.BotAPI, lifeService LifeService) *TelegramLifeCommander {
	return &TelegramLifeCommander{
		bot:         bot,
		lifeService: lifeService,
	}
}

func (telegramLifeCommander *TelegramLifeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case ListCommand:
		telegramLifeCommander.CallbackList(callback, callbackPath)
	default:
		telegramLifeCommander.sendError(callback.Message, "incorrect callback")
	}
}

func (telegramLifeCommander *TelegramLifeCommander) HandleCommand(inputMessage *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case HelpCommand:
		telegramLifeCommander.Help(inputMessage)
	case GetCommand:
		telegramLifeCommander.Get(inputMessage)
	case ListCommand:
		telegramLifeCommander.List(inputMessage)
	case DeleteCommand:
		telegramLifeCommander.Delete(inputMessage)
	case NewCommand:
		telegramLifeCommander.New(inputMessage)
	case EditCommand:
		telegramLifeCommander.Edit(inputMessage)
	default:
		telegramLifeCommander.sendError(inputMessage, "unknown command, use /help for get commands list")
	}
}

func (telegramLifeCommander *TelegramLifeCommander) sendError(inputMessage *tgbotapi.Message, errorMessage string) {
	telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "ERROR: "+errorMessage))
}
