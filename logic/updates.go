package logic

import (
	"context"
	"errors"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (l *Logic) HandleUpdate(ctx context.Context, upd tgbotapi.Update) error {
	if upd.Message != nil {
		chatId := int(upd.Message.Chat.ID)

		u, err := l.us.GetUser(ctx, chatId)
		if errors.Is(err, peacefulroad.ErrNotFound) {
			err := l.tg.SendWelcomeMessage(chatId, "http://test.dev:8080/login")
			if err != nil {
				fmt.Println(err)
			}
		}
		if err != nil {
			return fmt.Errorf("cannot get user: %w", err)
		}

		fmt.Println(u)
	}
	// if upd.CallbackQuery != nil {
	// 	if upd.CallbackQuery.Data == "show-status" {

	// 	}
	// }
	fmt.Println("received update", upd)
	return nil
}
