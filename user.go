package peacefulroad

import (
	"context"

	"golang.org/x/oauth2"
)

type User struct {
	ChatId      int
	Token       *oauth2.Token
	DisplayName string
	PersonID    int
}

type UserService interface {
	GetUser(ctx context.Context, chatId int) (User, error)
	SaveUser(ctx context.Context, u User) error
}
