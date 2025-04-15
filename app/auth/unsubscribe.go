package auth

import (
	"context"
	"fmt"

	// "mazza/app/utils"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"
	"mazza/firebase"
	"mazza/inits"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// Delete the current company and unsubscribe all associated users
func Unsubscribe(ctx *context.Context, client *generated.Client, activeUser *generated.User, activeCompany *generated.Company) error {
	// client := ent.FromContext(*ctx)
	if activeUser == nil || activeCompany == nil {
		return fmt.Errorf("unauthorized")
	}

	// Check if user is an admin of the active company
	isAdmin, err := activeUser.QueryAssignedRoles().
		Where(
			userrole.HasCompanyWith(company.ID(activeCompany.ID)),
			userrole.RoleEQ(userrole.RoleADMIN)).
		Exist(*ctx)
	if err != nil || !isAdmin {
		fmt.Println("err:", err)
		return fmt.Errorf("unauthorized")
	}

	allUsers, err := activeCompany.QueryUsers().All(*ctx)
	if err != nil || allUsers == nil {
		fmt.Println("err:", err)
		return fmt.Errorf("unauthorized")
	}

	var firebaseUIDs []string
	for _, user := range allUsers {
		firebaseUIDs = append(firebaseUIDs, user.FirebaseUID)
	}

	// Initiate transaction
	tx, err := client.Tx(*ctx)
	if err != nil {
		fmt.Println("tx err:", err)
		return fmt.Errorf("an error occurred")
	}

	_, err = tx.User.Delete().Where(
		user.HasCompanyWith(company.ID(activeCompany.ID)),
	).Exec(*ctx)
	if err != nil {
		fmt.Println("err:", err)
		tx.Rollback()
		return fmt.Errorf("an error occurred")
	}

	err = tx.Company.DeleteOneID(activeCompany.ID).Exec(*ctx)
	if err != nil {
		fmt.Println("err:", err)
		tx.Rollback()
		return fmt.Errorf("an error occurred")
	}

	// Delete all stored files of this company
	go firebase.DeleteAllFilesFromCompany(activeCompany.ID)
	go firebase.DeleteAllUsersFromCompany(activeCompany.ID, firebaseUIDs)

	err = tx.Commit()
	if err != nil {
		fmt.Println("err:", err)
		tx.Rollback()
		return fmt.Errorf("an error occurred")
	}

	// Notify user and company admin
	var expoPushTokens []expo.ExponentPushToken
	for _, user := range allUsers {
		if user.ExpoPushToken != nil && user.ID != activeUser.ID {
			expoPushTokens = append(expoPushTokens, expo.ExponentPushToken(*user.ExpoPushToken))
		}
	}

	var messages []expo.PushMessage
	if len(expoPushTokens) > 0 {
		messages = append(messages, expo.PushMessage{
			To:    expoPushTokens,
			Title: "Utilizador Excluído",
			Body:  fmt.Sprintf("A conta de %s foi excluída da empresa %s", activeUser.Name, activeCompany.Name),
			Sound: "default",
			Priority: expo.HighPriority,
		})
	}
	if activeUser.ExpoPushToken != nil {
		tokens := []expo.ExponentPushToken{
			expo.ExponentPushToken(*activeUser.ExpoPushToken),
		}
		messages = append(messages, expo.PushMessage{
			To:    tokens,
			Title: "Utilizador Excluído",
			Body:  fmt.Sprintf("A tua conta foi excluída da empresa %s", activeCompany.Name),
			Sound: "default",
			Priority: expo.HighPriority,
		})
	}

	if len(messages) > 0 {
		_, _ = inits.ExpoClient.PublishMultiple(messages)
	}
	
	return nil

	// companyGroup := []models.Company{company}
	// inits.DB.Raw(`SELECT * FROM users WHERE id = ? AND deleted_at IS NULL`, company.ID).Find(&allUsers)

	// // Check if there are other admin of the company
	// var numberOfAdmins int
	// for i := range allUsers {
	// 	if utils.Includes(models.UserRoles.Admin, allUsers[i].Roles) {
	// 		numberOfAdmins += 1
	// 	}
	// }

	// 	// If user is the only admin of the company...
	// 	if numberOfAdmins > 1 {
	// 		// Retrieve all daughter companies
	// 		var daughterCompanies []models.Company
	// 		inits.DB.Raw("SELECT * FROM companies WHERE parent_company_id = ? AND deleted_at IS NULL", company.ID).Find(&daughterCompanies)
	// 		companyGroup = append(companyGroup, daughterCompanies...)
	// 	}

	// 	for i := range companyGroup {
	// 		var companyIDs []int
	// 		if companyGroup[i].ID != company.ID {
	// 			companyIDs = append(companyIDs, int(companyGroup[i].ID))
	// 		}

	// 		if len(companyIDs) > 0 {
	// 			inits.DB.Raw(`SELECT * FROM users WHERE id = ? AND deleted_at IS NULL`, companyGroup[i]).Find(&allUsers)
	// 		}

	// 		// Retrieve all users associated with all companies from the current group
	// 		for _, user := range allUsers {
	// 			contacts[int(user.ID)] = contact{
	// 				ID:          user.ID,
	// 				CompanyName: company.Name,
	// 				Name:        user.Name,
	// 				Email:       user.Username,
	// 				IsAdmin:     utils.Includes(models.UserRoles.Admin, user.Roles),
	// 			}
	// 		}

	// 		// Set delete flag on all companies from the group
	// 		inits.DB.Delete(&companyGroup)

	// 		// Set delete flag on all users from the group
	// 		inits.DB.Delete(&allUsers)
	// 	}
	// } else {
	// 	// If the current user in not an admin...
	// 	// Set delete flag on the user
	// 	inits.DB.Delete(&user)

	// 	contacts[int(user.ID)] = contact{
	// 		ID:          user.ID,
	// 		CompanyName: company.Name,
	// 		Name:        user.Name,
	// 		Email:       user.Username,
	// 		IsAdmin:     utils.Includes(models.UserRoles.Admin, user.Roles),
	// 	}
	// }

	// // Send farewell emails
	// for _, c := range contacts {
	// 	// If user is not the only admin, only remove him and return
	// 	sendFarewellEmail(&c)
	// }

	// // c.JSON(http.StatusOK, gin.H{"message": "User was successfully unsubscribed"})
	// c.JSON(http.StatusOK, gin.H{"admin": isAdmin, "contacts": contacts})
}
