package firebase

import (
	"context"
	"fmt"

	"github.com/JakubC-projects/pacebot"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Store) GetUser(ctx context.Context, chatId int) (pacebot.User, error) {
	var res pacebot.User
	doc, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(chatId)).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return pacebot.User{}, pacebot.ErrNotFound
	}
	if err != nil {
		return res, fmt.Errorf("cannot fetch user from the database: %w", err)
	}
	err = doc.DataTo(&res)
	return res, err
}

func (s *Store) GetAllUsers(ctx context.Context) ([]pacebot.User, error) {
	var res []pacebot.User
	iter := s.client.Collection(userCollectionName).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot get document: %w", err)
		}
		var user pacebot.User
		if err := doc.DataTo(&user); err != nil {
			return nil, fmt.Errorf("cannot marshal document: %w", err)
		}
		res = append(res, user)
	}

	return res, nil
}

func (s *Store) SaveUser(ctx context.Context, user pacebot.User) error {
	_, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(user.ChatId)).Set(ctx, user)
	return err
}

func (s *Store) DeleteUser(ctx context.Context, chatId int) error {
	_, err := s.client.Collection(userCollectionName).Doc(fmt.Sprint(chatId)).Delete(ctx)
	return err
}
