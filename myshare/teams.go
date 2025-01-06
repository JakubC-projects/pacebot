package myshare

import (
	"context"

	"github.com/JakubC-projects/pacebot"
)

type Org struct {
	Id int `json:"id"`
}

func (c *Client) GetOrgs(ctx context.Context, u pacebot.User) ([]Org, error) {
	var res response[[]Org]

	err := c.get(ctx, u.Token, "/Profile/Organisations", &res)
	return res.Data, err
}
