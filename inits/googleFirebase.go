package inits

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	FCM             *messaging.Client
	FirebaseStorage *storage.BucketHandle
	Firestore       *firestore.Client
)

func InitFirebase() error {
	ctx := context.Background()
	config := &firebase.Config{StorageBucket: os.Getenv("FIREBASE_BUCKET")}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CONFIG_FILE"))

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal("Firebase initialization error:", err)
	}

	// Initialize Messaging client
	FCM, err = app.Messaging(ctx)
	if err != nil {
		log.Fatal("Firebase storage initialization error:", err)
	}

	// Initialize Storage client
	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatal("Firebase storage initialization error:", err)
	}

	FirebaseStorage, err = client.DefaultBucket()
	if err != nil {
		log.Fatal("Firebase storage initialization error:", err)
	}

	// Initialize Firestore client
	Firestore, err = app.Firestore(ctx)

	return err
}
