package firebase

import (
	"context"
	"fmt"
	"mazza/ent/generated/userrole"
	"mazza/inits"
	"strconv"
)

/* Create a new entry in the users collection in the Firestore Database */
func CreateUserEntry(ctx context.Context, companyID int, firebaseUID string, isActive bool, role userrole.Role) error {
	inits.Firestore.Collection("users")
	result, err := inits.Firestore.Collection("users").Doc(firebaseUID).Set(ctx, map[string]any{
		"companyId": strconv.Itoa(companyID),
		"active":    isActive,
		"role":      role,
	})
	fmt.Println("CreateUserEntry err:", err, "\nresult:", result)
	return err
}
