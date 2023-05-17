package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
)

type Telegram struct {
}

func New() *Telegram {
	var t Telegram
	return &t
}

func (t *Telegram) SendMessage(message string) {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Some error occured. Err: %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	msg := tgbotapi.NewMessage(907685216, message)
	_, err = bot.Send(msg)
	if err != nil {
		fmt.Println("Send message error: ", err)
	}
}
