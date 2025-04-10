package firebase

import (
	"context"
	"fmt"
	"log"
	"mazza/inits"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// Deletes all stored files from a given company
func DeleteAllFilesFromCompany(companyID int) error {
	// Set timeout for the delete operation
	var ctx = context.Background()
	deleteCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	folderPath := strconv.Itoa(companyID) + "/"

	// List all objects in the folder
	iter := inits.FirebaseStorage.Objects(deleteCtx, &storage.Query{Prefix: folderPath})

	// Delete all objects in the folder
	for {
		objAttrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return err
		}

		if err := inits.FirebaseStorage.Object(objAttrs.Name).Delete(deleteCtx); err != nil {
			log.Printf("error deleting object %s: %v", objAttrs.Name, err)
		}
	}
	// err := storageHandler.Delete(deleteCtx)
	// fmt.Printf("DeleteAllFromCompany successfully deleted folder: %s\n", folderPath)
	return nil
}

func DeleteAllUsersFromCompany(companyID int, firebaseUIDs []string) error {
	var ctx = context.Background()
	deleteCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	compID := strconv.Itoa(companyID)
	iter := inits.Firestore.Collection("users").Where("companyId", "==", compID).Documents(deleteCtx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("DeleteAllUsersFromCompany err:", err)
			return err
		}

		_, err = doc.Ref.Delete(ctx)
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	_, err := inits.Auth.DeleteUsers(deleteCtx, firebaseUIDs)
	if err != nil {
		fmt.Println("err:", err)
	}

	return nil
}
