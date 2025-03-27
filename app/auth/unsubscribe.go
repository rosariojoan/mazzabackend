package auth

import (
	"context"
	"fmt"

	// "mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/userrole"
	"mazza/ent/utils"
)

type contact struct {
	ID          uint
	CompanyName string
	Name        string
	Email       string
	IsAdmin     bool
}

func Unsubscribe(ctx *context.Context) error {
	client := ent.FromContext(*ctx)
	currentUser, company := utils.GetSession(ctx)
	if currentUser == nil || company == nil {
		return fmt.Errorf("unauthorized")
	}

	fmt.Println("ctx:", currentUser, company, client)
	// contacts := map[int]contact{}
	// fmt.Println("user:", user, company)

	// Check if the user is admin
	isAdmin, err := currentUser.QueryAssignedRoles().Where(userrole.RoleEQ(userrole.RoleADMIN)).Exist(*ctx)
	if err != nil || !isAdmin {
		fmt.Println("err:", err)
		return fmt.Errorf("unauthorized")
	}

	allUsers, err := currentUser.QueryAssignedRoles().WithCompany().QueryUser().All(*ctx)
	if err != nil || allUsers == nil {
		fmt.Println("err:", err)
		return fmt.Errorf("unauthorized")
	}
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

	return nil
}

// Send a farewell email to the deleted user.
func sendFarewellEmail(contact *contact) {
	// ...
	fmt.Println("sent:", contact.ID)
}
