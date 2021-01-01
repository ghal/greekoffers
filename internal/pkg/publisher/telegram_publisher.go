package publisher

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"strconv"
)

// TelegramPublisher struct.
type TelegramPublisher struct{}

// NewTelegramPublisher returns a TelegramPublisher.
func NewTelegramPublisher() *TelegramPublisher {
	return &TelegramPublisher{}
}

// Publish publishes the scraped item to telegram.
func (dv *TelegramPublisher) Publish(item Item) {
	telegramToken, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		fmt.Errorf("TELEGRAM_TOKEN not found")
	}

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	bot.Debug = true

	chatIDStr, ok := os.LookupEnv("TELEGRAM_CHAT_ID")
	if !ok {
		fmt.Errorf("CHAT_ID not found")
	}
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	msg := tgbotapi.NewMessage(
		int64(chatID),
		item.Title+": \n"+
			item.URL+"\n",
	)

	_, err = bot.Send(msg)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
