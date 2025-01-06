package myshare

import (
	"context"
	"errors"
	"fmt"

	"github.com/JakubC-projects/pacebot"
	"golang.org/x/oauth2"
)

type myshareStatus struct {
	TransactionsAmount float64         `json:"transactionsAmount"`
	PercentageValue    float64         `json:"percentageValue"`
	Targets            []myshareTarget `json:"targets"`
}

type myshareTarget struct {
	Currency    string  `json:"currency"`
	TotalAmount float64 `json:"totalAmount"`
}

func (c *Client) GetStatus(ctx context.Context, token *oauth2.Token, u pacebot.User) (pacebot.Status, error) {
	path := fmt.Sprintf("/TargetStatus/%d/Member/%d", u.ClubId, u.PersonID)

	var res response[myshareStatus]

	err := c.get(ctx, token, path, &res)

	if len(res.Data.Targets) < 1 {
		return pacebot.Status{}, errors.New("no target specified")
	}

	currentTarget := res.Data.Targets[0]

	status := pacebot.Status{
		TransactionsAmount: res.Data.TransactionsAmount,
		PercentageValue:    res.Data.PercentageValue,
		Target:             currentTarget.TotalAmount,
		Currency:           currentTarget.Currency,
	}
	return status, err
}
