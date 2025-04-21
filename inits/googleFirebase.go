package inits

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"google.golang.org/api/option"
)

var (
	Auth            *auth.Client
	FCM             *messaging.Client
	FirebaseStorage *storage.BucketHandle
	Firestore       *firestore.Client
	// Use this variable to send push notification to the client
	ExpoClient *expo.PushClient
)

type ServiceAccount struct {
	Type              string `json:"type"`
	ProjectID         string `json:"project_id"`
	PrivateKeyID      string `json:"private_key_id"`
	ProjectKey        string `json:"private_key"`
	ClientEmail       string `json:"client_email"`
	ClientID          string `json:"client_id"`
	AuthURI           string `json:"auth_uri"`
	TokenURI          string `json:"token_uri"`
	AuthProvider_x509 string `json:"auth_provider_x509_cert_url"`
	Client_x509       string `json:"client_x509_cert_url"`
	UniverseDomain    string `json:"universe_domain"`
}

func InitFirebase() error {
	ctx := context.Background()
	serviceAccount := ServiceAccount{
		Type:              os.Getenv("SERVICE_ACCOUNT_TYPE"),
		ProjectID:         os.Getenv("SERVICE_ACCOUNT_PROJECT_ID"),
		PrivateKeyID:      os.Getenv("SERVICE_ACCOUNT_PRIVATE_KEY_ID"),
		ProjectKey:        os.Getenv("SERVICE_ACCOUNT_PROJECT_KEY"),
		ClientEmail:       os.Getenv("SERVICE_ACCOUNT_CLIENT_EMAIL"),
		ClientID:          os.Getenv("SERVICE_ACCOUNT_CLIENT_ID"),
		AuthURI:           os.Getenv("SERVICE_ACCOUNT_AUTH_URI"),
		TokenURI:          os.Getenv("SERVICE_ACCOUNT_TOKEN_URI"),
		AuthProvider_x509: os.Getenv("SERVICE_ACCOUNT_AUTH_PROVIDER_X509"),
		Client_x509:       os.Getenv("SERVICE_ACCOUNT_CLIENT_X509"),
		UniverseDomain:    os.Getenv("SERVICE_ACCOUNT_UNIVERSE_DOMAIN"),
	}

	serviceAccountJSON, err := json.Marshal(serviceAccount)
	if err != nil {
		log.Fatalln("unable to start firebase:", err)
	}

	config := &firebase.Config{StorageBucket: os.Getenv("FIREBASE_BUCKET")}
	opt := option.WithCredentialsJSON(serviceAccountJSON)
	// _ = serviceAccountJSON
	// opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CONFIG_FILE"))

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal("Firebase initialization error:", err)
	}

	Auth, err = app.Auth(ctx)
	if err != nil {
		log.Fatal("Firebase auth initialization error:", err)
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
	if err != nil {
		log.Fatal("Firebase storage initialization error:", err)
	}

	// Expo client
	ExpoClient = expo.NewPushClient(nil)

	return err
}
