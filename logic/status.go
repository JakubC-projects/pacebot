package logic

import (
	"context"
	"fmt"
	"time"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

func (l *Logic) getStatusMessage(ctx context.Context, user peacefulroad.User) (peacefulroad.StatusMessage, error) {
	status, err := l.ms.GetStatus(ctx, user.Token, user)
	if err != nil {
		return peacefulroad.StatusMessage{}, fmt.Errorf("cannot get user status: %w", err)
	}

	statusMessage := peacefulroad.StatusMessage{
		CurrentStatus: status.TransactionsAmount,
		SeasonTarget:  status.Target,
		Currency:      status.Currency,
		WeekTarget:    getStatusForWeek(time.Now()),

		RegisterURL:   "https://app.myshare.today/registration",
		DonateURL:     "https://donationbuk.no",
		LogoutURL:     l.auth.LogoutEndpoint(user.ChatId),
		ShowNotifyAll: user.IsAdmin,
	}

	return statusMessage, nil
}
