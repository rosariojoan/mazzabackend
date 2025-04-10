package firebase

import (
	"context"
	"fmt"
	"mazza/inits"
	"time"
)

// Delete an item from Firestorage
func DeleteItem(fileURI string) error {
	// Reference to the file
	fileRef := inits.FirebaseStorage.Object(fileURI)

	// Set timeout for the delete operation
	var ctx = context.Background()
	deleteCtx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	// Delete the file
	if err := fileRef.Delete(deleteCtx); err != nil {
		fmt.Println("v fileURI:", fileURI)
		fmt.Println("firebase DeleteItem err:", err)
		return err
	}

	return nil
}
