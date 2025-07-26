package notifications

import (
	"context"
	"fmt"
	"mazza/app/notifications/expo"
	"mazza/inits"
	"time"
)

// Send Expo push notifications. The operation is cancelled if it exceeds ctxDeadline of time
func SendPushNotification(ctxDeadline time.Duration, msgs []*expo.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxDeadline)
	defer cancel()

	_, err := inits.ExpoClient.Publish(ctx, msgs)

	if err != nil {
		fmt.Println("SendNotification:", err)
		// utils.Pp(msgs)
	}
	// else {
	// 	fmt.Println("SendNotification:", &res)
	// }
}
