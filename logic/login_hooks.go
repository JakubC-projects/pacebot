package logic

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

func (a *Logic) postLoginHook(ctx context.Context, user peacefulroad.User) error {
	err := a.us.SaveUser(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}

	statusMessage, err := a.getStatusMessage(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot get status message: %w", err)
	}

	err = a.tg.SendStatusMessage(user.ChatId, statusMessage)
	if err != nil {
		return fmt.Errorf("cannot send status message: %w", err)
	}
	return nil
}

func (a *Logic) postLogoutHook(ctx context.Context, user peacefulroad.User) error {
	err := a.us.DeleteUser(ctx, user.ChatId)
	if err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}

	err = a.tg.SendWelcomeMessage(user.ChatId, a.auth.LoginEndpoint(user.ChatId))
	if err != nil {
		return fmt.Errorf("cannot send status message: %w", err)
	}
	return nil
}
