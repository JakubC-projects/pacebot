package logic

import (
	"context"
	"errors"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (l *Logic) HandleUpdate(ctx context.Context, upd tgbotapi.Update) error {
	var chatId = getChatId(upd)

	user, err := l.us.GetUser(ctx, chatId)
	if errors.Is(err, peacefulroad.ErrNotFound) {
		err := l.tg.SendWelcomeMessage(chatId, l.auth.LoginEndpoint(chatId))
		if err != nil {
			fmt.Println(err)
		}
	}
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	statusMessage, err := l.getStatusMessage(ctx, user)
	if err != nil {
		return err
	}

	if upd.CallbackQuery != nil {
		if upd.CallbackQuery.Data == "show-status" {
			err = l.tg.EditStatusMessage(chatId, upd.CallbackQuery.Message.MessageID, statusMessage)
			if err != nil {
				return err
			}

			return nil
		}
	}
	err = l.tg.SendStatusMessage(chatId, statusMessage)
	if err != nil {
		return err
	}

	return nil
}

func getChatId(upd tgbotapi.Update) int {
	if upd.Message != nil {
		return int(upd.Message.Chat.ID)
	}
	if upd.CallbackQuery != nil {
		return int(upd.CallbackQuery.From.ID)
	}
	return 0
}
