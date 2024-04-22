package myshare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var (
	baseUrl = os.Getenv("MYSHARE_BASE_URL")
	clubId  = os.Getenv("MYSHARE_CLUB_ID")
)

func TestGetStatus(t *testing.T) {
	tokenRaw := ""
	personId := 41838
	token := &oauth2.Token{AccessToken: tokenRaw, TokenType: "Bearer"}

	client := NewClient(baseUrl, clubId)

	status, err := client.GetStatus(context.Background(), token, personId)

	assert.NoError(t, err)

	fmt.Println(status)
}
