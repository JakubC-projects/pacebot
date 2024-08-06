package myshare

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type response[T any] struct {
	StatusCode int `json:"statusCode"`
	Data       T   `json:"data"`
}

func (c *Client) get(ctx context.Context, token *oauth2.Token, path string, result any) error {
	url := c.baseUrl + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed: invalid status (%d): %s", resp.StatusCode, body)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("cannot unmarshal result: %w", err)
	}

	return nil
}
