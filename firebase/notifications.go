package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mazza/inits"

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
	resp, err := inits.FCM.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: *title,
			Body:  "",
		},
		Data:  data,
		Tokens: FCMTokens,
		// Topic: "test-topic",
		// Token: "ezKh5hXeTeuNtMh4CXtIHj:APA91bEA21Qvt9knhB_iKydYYjJO5Ks48dNEcUGfjSGPzNMQvGOAQ21NmlsJDDyN5QetWP9np8JJNTMW4pkzMTgn3Mz5S17FtBqtKBouOP73DWsGIiK1I5f1LLcANMgDXe4AqvtCALr-",
		// Token: "e67v_ojgQ7m0zrCrlgMM9S:APA91bE3wxbp6NJcIpfB0ETxUkZ3xv-Ig0Pu9oW1om7xDKMQMBoWg_VaHRPzxuZxwJvNmFzez4ikQQTKr3KdS1yCKU9rr3da2OQxhS0bixmFuKEvX8nth6AP1Cu2ayWCDcJ5kIzNBcqV",
	})
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
		return
	}
	fmt.Println("Nofit sent successful:", resp.Responses)
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
