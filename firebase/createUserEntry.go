package firebase

import (
	"context"
	"mazza/ent/generated/userrole"
	"mazza/inits"
	"strconv"
)

/* Create a new entry in the users collection in the Firestore Database */
func CreateUserEntry(ctx context.Context, companyID int, firebaseUID string, isActive bool, role userrole.Role) error {
	_, err := inits.Firestore.Collection("users").Doc(firebaseUID).Set(ctx, map[string]any{
		"companyId": strconv.Itoa(companyID),
		"active":    isActive,
		"role":      role,
	})
	return err
}
