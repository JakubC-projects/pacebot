package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	peacefulroad "github.com/JakubC-projects/peaceful-road"
	_ "google.golang.org/api/option"
)

const userCollectionName = "users"

type Store struct {
	client *firestore.Client
}

var _ peacefulroad.UserService = (*Store)(nil)

func NewStore(projectId string) *Store {
	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		panic(fmt.Errorf("cannot create firebase app: %w", err))
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(fmt.Errorf("cannot create firestore client: %w", err))
	}
	return &Store{client}
}
