package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) SendStatusMessage(chatId int) error {
	text := "Status Message"
	_, err := s.bot.Send(tgbotapi.NewMessage(int64(chatId), text))

	return err
}
