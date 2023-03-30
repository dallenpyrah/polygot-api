package interfaces

import (
	firebaseStorage "firebase.google.com/go/storage"
)

type FirebaseInitializer interface {
	CreateClient() (*firebaseStorage.Client, error)
}
