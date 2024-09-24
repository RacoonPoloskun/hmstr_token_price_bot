package internal

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
	"strconv"
)

type TelegramBot struct {
	bot       *telego.Bot
	channelId int64
}

func NewTelegramBot() *TelegramBot {
	channelId, err := strconv.ParseInt(os.Getenv("TG_CHANNEL_ID"), 10, 64)
	if err != nil {
		log.Println("Инициализация бота:", err)
		os.Exit(1)
	}

	bot, err := telego.NewBot(os.Getenv("TG_BOT_TOKEN"), telego.WithDefaultDebugLogger())
	if err != nil {
		log.Println("Инициализация бота:", err)
		os.Exit(1)
	}

	return &TelegramBot{
		bot:       bot,
		channelId: channelId,
	}
}

func (tb *TelegramBot) SendMessageToChannel(message string) error {
	_, err := tb.bot.SendMessage(tu.Message(tu.ID(tb.channelId), message))
	return err
}
