package firebase

import (
	"context"
	"fmt"
	"mazza/app/notifications"
	"mazza/app/notifications/expo"
	"mazza/ent/generated/userrole"
	"mazza/inits"
	"time"

	"cloud.google.com/go/firestore"
)

func UpdateUser(companyID int, companyName string, firebaseUID string, isActive *bool, role userrole.Role, expoPushToken *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var updates = []firestore.Update{{
		Path:  "role",
		Value: role,
	}}
	if isActive != nil {
		updates = append(updates, firestore.Update{
			Path:  "active",
			Value: isActive,
		})
	}

	_, err := inits.Firestore.Collection("users").Doc(firebaseUID).Update(ctx, updates)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	if isActive == nil || expoPushToken == nil {
		return nil
	}

	// Notify user
	var title string
	var body string
	var data map[string]string

	var status string
	if *isActive {
		status = "active"
		title = "A tua conta foi activada"
		body = "Agora podes trabalhar na " + companyName + " de acordo com as tuas funções."
	} else {
		status = "inactive"
		title = "A tua conta foi desactivada"
		body = "Continuas tendo acesso ao Mazza mas o teu acesso à " + companyName + " está limitado."
	}
	data = map[string]string{
		"type":    AlertType.UserActiveStatusChange,
		"status":  status,
		"company": companyName,
	}
	// // If user role was changed
	// data = map[string]string{
	// 	"type":    AlertType.UserRoleChange,
	// 	"role":    role.String(),
	// 	"company": companyName,
	// }

	token := []*expo.Token{expo.MustParseToken(*expoPushToken)}
	go notifications.SendPushNotification(15*time.Second, []*expo.Message{{
		To:       token,
		Title:    title,
		Body:     body,
		Data:     data,
		Sound:    "default",
		Priority: expo.HighPriority,
	}})

	return nil
}
