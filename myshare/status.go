package myshare

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

func (c *Client) GetStatus(ctx context.Context, u peacefulroad.User) (peacefulroad.Status, error) {
	path := fmt.Sprintf("/TargetStatus/%d/Member/%d", peacefulroad.BukUstronClub, u.PersonID)

	var res response[peacefulroad.Status]

	err := c.get(ctx, u.Token, path, &res)
	return res.Data, err
}
