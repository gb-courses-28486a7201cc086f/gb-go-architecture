package notification

import (
	"fmt"
	"gb-go-architecture/lesson-2/shop_new/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	chatID int64
	tgBot  *tgbotapi.BotAPI
}

func (s *TelegramBot) SendOrderCreated(order *models.Order) error {
	text := fmt.Sprintf("new order %d\n\nphone: %s", order.ID, order.CustomerPhone)

	fmt.Printf("message to TG chat %d about order %d\n", s.chatID, order.ID)
	msg := tgbotapi.NewMessage(s.chatID, text)

	_, err := s.tgBot.Send(msg)
	return err
}

func NewTelegramBot(token string, chatID int64) (*TelegramBot, error) {
	tgBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		chatID: chatID,
		tgBot:  tgBot,
	}, nil
}
