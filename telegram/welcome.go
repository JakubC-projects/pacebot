package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) SendWelcomeMessage(chatId int, loginUrl string) error {
	text := "Zaloguj się aby rozpocząć!"

	buttons := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{{Text: "Login", URL: &loginUrl}},
		},
	}

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      int64(chatId),
			ReplyMarkup: buttons,
		},
		Text: text,
	}

	_, err := s.bot.Send(msg)

	return err
}
