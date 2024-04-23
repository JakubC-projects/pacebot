package telegram

import (
	"fmt"
	"time"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

func (s *Service) SendStatusMessage(chatId int, content peacefulroad.StatusMessage) error {
	text, buttons := s.getStatusMessage(content)

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      int64(chatId),
			ReplyMarkup: buttons,
		},
		Text:      text,
		ParseMode: tgbotapi.ModeHTML,
	}

	_, err := s.bot.Send(msg)

	return err
}

func (s *Service) EditStatusMessage(chatId int, messageId int, content peacefulroad.StatusMessage) error {
	text, buttons := s.getStatusMessage(content)

	msg := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      int64(chatId),
			MessageID:   messageId,
			ReplyMarkup: &buttons,
		},
		Text:      text,
		ParseMode: tgbotapi.ModeHTML,
	}

	_, err := s.bot.Send(msg)

	return err
}

func (s *Service) getStatusMessage(content peacefulroad.StatusMessage) (string, tgbotapi.InlineKeyboardMarkup) {
	userPercent := content.CurrentStatus / content.SeasonTarget * 100

	missingAmount := (content.WeekTarget - userPercent) * content.SeasonTarget / 100

	statusEmoji := "🟢"
	statusMessage := ""

	if missingAmount > 0 {
		statusEmoji = "🔴"
		statusMessage = fmt.Sprintf("Brakuje ci: <b>%.2f</b> %s\n", missingAmount, content.Currency)
	}

	text := fmt.Sprintf(`
Cel na ten tydzień: <b>%.2f%%</b>
Twój Status: <b>%.2f%%</b> %s (%.2f / %.2f %s)
%s
<a href="%s">Zapisz się na Dugnad!</a>
<a href="%s">Wpłać na MyShare!</a>

Dane z: %s
`,
		content.WeekTarget,
		userPercent,
		statusEmoji,
		content.CurrentStatus,
		content.SeasonTarget,
		content.Currency,
		statusMessage,
		content.RegisterURL,
		content.DonateURL,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	buttons := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{{Text: "Odśwież", CallbackData: lo.ToPtr("show-status")}},
			{{Text: "Wyloguj się", URL: lo.ToPtr(content.LogoutURL)}},
		}}

	return text, buttons
}
