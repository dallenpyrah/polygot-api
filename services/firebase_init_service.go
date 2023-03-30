package services

import (
	"context"
	firebaseStorage "firebase.google.com/go/storage"
	"fmt"
	"polygot-api/providers"
	"time"
)

type FirebaseInitService struct {
	firebaseStorageProvider providers.FirebaseStorageProvider
}

func (f FirebaseInitService) CreateClient() (*firebaseStorage.Client, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	client, err := f.firebaseStorageProvider.CreateFirebaseStorageClient(context)

	if err != nil {
		return nil, fmt.Errorf("error initializing storage client: %v", err)
	}

	return client, nil
}
