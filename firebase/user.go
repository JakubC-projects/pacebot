package firebase

import (
	"context"
	"fmt"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Store) GetUser(ctx context.Context, chatId int) (peacefulroad.User, error) {
	var res peacefulroad.User
	doc, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(chatId)).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return peacefulroad.User{}, peacefulroad.ErrNotFound
	}
	if err != nil {
		return res, fmt.Errorf("cannot fetch user from the database: %w", err)
	}
	err = doc.DataTo(&res)
	return res, err
}

func (s *Store) SaveUser(ctx context.Context, user peacefulroad.User) error {
	_, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(user.ChatId)).Set(ctx, user)
	return err
}

func (s *Store) DeleteUser(ctx context.Context, chatId int) error {
	_, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(chatId)).Delete(ctx)
	return err
}
