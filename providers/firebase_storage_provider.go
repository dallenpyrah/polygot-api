package providers

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

type FirebaseStorageProvider struct {
}

func (f FirebaseStorageProvider) CreateFirebaseStorageClient(ctx context.Context) (*storage.Client, error) {
	credentialFilePath := os.Getenv("FIREBASE_CREDENTIALS")

	if credentialFilePath == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIALS environment variable is not set")
	}

	clientOptions := option.WithCredentialsFile(credentialFilePath)
	app, err := firebase.NewApp(ctx, nil, clientOptions)

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing storage client: %v", err)
	}

	return client, nil
}
