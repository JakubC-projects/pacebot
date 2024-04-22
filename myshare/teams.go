package myshare

import (
	"context"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

type Org struct {
}

func (c *Client) GetOrgs(ctx context.Context, u peacefulroad.User) ([]Org, error) {
	var res response[[]Org]

	err := c.get(ctx, u.Token, "/Profile/Organisations", &res)
	return res.Data, err
}
