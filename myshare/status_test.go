package myshare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/JakubC-projects/pacebot"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var (
	baseUrl = os.Getenv("MYSHARE_BASE_URL")
)

func TestGetStatus(t *testing.T) {
	tokenRaw := ""
	u := pacebot.User{
		PersonID: 41838,
		ClubId:   3982,
	}
	token := &oauth2.Token{AccessToken: tokenRaw, TokenType: "Bearer"}

	client := NewClient(baseUrl)

	status, err := client.GetStatus(context.Background(), token, u)

	assert.NoError(t, err)

	fmt.Println(status)
}
