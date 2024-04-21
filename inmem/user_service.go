package inmem

import (
	"context"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

type UserService struct {
	users []peacefulroad.User
}

func NewUserService() *UserService {
	return &UserService{}
}

var _ peacefulroad.UserService = (*UserService)(nil)

func (us *UserService) GetUser(ctx context.Context, chatId int) (peacefulroad.User, error) {
	for _, u := range us.users {
		if u.ChatId == chatId {
			return u, nil
		}
	}
	return peacefulroad.User{}, peacefulroad.ErrNotFound
}

func (us *UserService) SaveUser(ctx context.Context, user peacefulroad.User) error {
	for i, u := range us.users {
		if u.ChatId == user.ChatId {
			us.users[i] = user
		}
	}
	us.users = append(us.users, user)
	return nil
}
