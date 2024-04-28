package logic

import (
	"context"
	"errors"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/sync/errgroup"
)

func (l *Logic) HandleUpdate(ctx context.Context, upd tgbotapi.Update) error {
	var chatId = getChatId(upd)

	user, err := l.us.GetUser(ctx, chatId)
	if errors.Is(err, peacefulroad.ErrNotFound) {
		err := l.tg.SendWelcomeMessage(chatId, l.auth.LoginEndpoint(chatId))
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}
	if err != nil {
		err := fmt.Errorf("cannot get user: %w", err)
		fmt.Println(err.Error())
		l.tg.SendErrorMessage(user.ChatId, err.Error())
		return err
	}

	err = l.ensureValidToken(ctx, &user)
	if err != nil {
		fmt.Println(err.Error())
		err := l.tg.SendWelcomeMessage(chatId, l.auth.LoginEndpoint(chatId))
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	if upd.CallbackQuery != nil {
		msgId := upd.CallbackQuery.Message.MessageID
		var err error
		switch upd.CallbackQuery.Data {
		case "notify-all":
			err = l.handleNotifyAll(ctx, user)
		case "show-status":
			err = l.handleUpdateStatus(ctx, user, msgId)
		}

		if err != nil {
			l.tg.SendErrorMessage(user.ChatId, err.Error())
		}

		return nil
	}

	// When in doubt send an update message
	statusMessage, err := l.getStatusMessage(ctx, user)
	if err != nil {
		return err
	}
	err = l.tg.SendStatusMessage(chatId, statusMessage)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logic) handleUpdateStatus(ctx context.Context, user peacefulroad.User, messageId int) error {
	statusMessage, err := l.getStatusMessage(ctx, user)
	if err != nil {
		return err
	}

	err = l.tg.EditStatusMessage(user.ChatId, messageId, statusMessage)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logic) handleNotifyAll(ctx context.Context, user peacefulroad.User) error {
	if !user.IsAdmin {
		return nil
	}

	allUsers, err := l.us.GetAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("cannot get all users: %w", err)
	}

	eg := errgroup.Group{}

	for _, user := range allUsers {
		user := user
		eg.Go(func() error {
			err := l.ensureValidToken(ctx, &user)
			if err != nil {
				fmt.Printf("cannot renew token for %s: %s\n", user.DisplayName, err)
				l.tg.SendWelcomeMessage(user.ChatId, l.auth.LoginEndpoint(user.ChatId))
				return nil
			}
			statusMessage, err := l.getStatusMessage(ctx, user)
			if err != nil {
				err = fmt.Errorf("cannot calculate status message for %s: %w", user.DisplayName, err)
				fmt.Println(err)
				return err
			}
			err = l.tg.SendStatusMessage(user.ChatId, statusMessage)
			if err != nil {
				err = fmt.Errorf("cannot send status message to %s: %w", user.DisplayName, err)
				fmt.Println(err)
				return err
			}
			return nil
		})
	}

	return eg.Wait()
}

func (l *Logic) ensureValidToken(ctx context.Context, user *peacefulroad.User) error {
	tok, err := l.auth.GetFreshToken(ctx, user.Token)
	if err != nil {
		return fmt.Errorf("cannot get fresh token: %w", err)
	}
	if tok.AccessToken != user.Token.AccessToken {
		user.Token = tok
		err := l.us.SaveUser(ctx, *user)
		if err != nil {
			return fmt.Errorf("cannot save user: %w", err)
		}
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
