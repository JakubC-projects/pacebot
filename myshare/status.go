package myshare

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"golang.org/x/oauth2"
)

func (c *Client) GetStatus(ctx context.Context, token *oauth2.Token, u peacefulroad.User) (peacefulroad.Status, error) {
	path := fmt.Sprintf("/TargetStatus/%d/Member/%d", u.ClubId, u.PersonID)

	var res response[peacefulroad.Status]

	err := c.get(ctx, token, path, &res)
	return res.Data, err
}
