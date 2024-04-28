package logic

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/JakubC-projects/peaceful-road/myshare"
	"github.com/samber/lo"
)

func (a *Logic) postLoginHook(ctx context.Context, user peacefulroad.User) error {
	orgs, err := a.ms.GetOrgs(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot get user orgs: %w", err)
	}
	orgIds := lo.Map(orgs, func(o myshare.Org, _ int) int { return o.Id })

	userClubs := lo.Intersect(orgIds, a.allowedClubIds)
	if len(userClubs) < 1 {
		return fmt.Errorf("not a user of supported club")
	}

	user.ClubId = userClubs[0]

	err = a.us.SaveUser(ctx, user)
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
