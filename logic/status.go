package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/JakubC-projects/pacebot"
)

func (l *Logic) getStatusMessage(ctx context.Context, user pacebot.User) (pacebot.StatusMessage, error) {
	status, err := l.ms.GetStatus(ctx, user.Token, user)
	if err != nil {
		return pacebot.StatusMessage{}, fmt.Errorf("cannot get user status: %w", err)
	}

	statusMessage := pacebot.StatusMessage{
		CurrentStatus:   status.TransactionsAmount,
		SeasonTarget:    status.Target,
		Currency:        status.Currency,
		MilestoneTarget: getStatusForNextMilestone(time.Now()),

		RegisterURL:   "https://app.myshare.today/registration",
		DonateURL:     "https://donationbuk.no",
		LogoutURL:     l.auth.LogoutEndpoint(user.ChatId),
		ShowNotifyAll: user.IsAdmin,
	}

	return statusMessage, nil
}
