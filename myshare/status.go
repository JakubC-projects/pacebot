package myshare

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"golang.org/x/oauth2"
)

func (c *Client) GetStatus(ctx context.Context, token *oauth2.Token, personId int) (peacefulroad.Status, error) {
	path := fmt.Sprintf("/TargetStatus/%s/Member/%d", c.clubId, personId)

	var res response[peacefulroad.Status]

	err := c.get(ctx, token, path, &res)
	return res.Data, err
}
