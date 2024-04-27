package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) SendErrorMessage(chatId int, errMsg string) error {

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: int64(chatId),
		},
		Text: fmt.Sprintf("Encountered an error: %s", errMsg),
	}

	_, err := s.bot.Send(msg)

	return err
}
