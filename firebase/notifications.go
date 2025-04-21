package firebase

import (
	"context"
	"encoding/json"

	"firebase.google.com/go/messaging"
)

// SendFCMMessage sends an FCM message using Firebase messaging client.
func SendFCMMessage(ctx context.Context, message *messaging.Message) error {
	// Get the FirebaseUtil singletion instance
	// firebaseUtil, err := inits.GetFirebaseUtil(ctx)
	// if err != nil {
	// 	return err
	// }

	// // Get the Messaging client
	// messagingClient := firebaseUtil.GetMessagingClient()

	// // Send the message
	// response, err := messagingClient.Send(ctx, message)
	// if err != nil {
	// 	return err
	// }

	// log.Println("Successfully sent FCM message:", response)
	return nil
}

/* Send data notification to a single user. */
func SendDataNotification(FCMTokens []string, title *string, data map[string]string) {
	return
}

// Convert struct to map to be sent sent as notification data
func GetMap(data any) (map[string]string, error) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	dataMap := map[string]string{
		"jsonData": string(dataJson),
	}
	return dataMap, nil
}
