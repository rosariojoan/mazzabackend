package utils

import (
	"bytes"
	"context"
	"io"
	"log"
	"mazza/inits"
	"strconv"

	"cloud.google.com/go/storage"
)

type FileProps struct {
	Bucket      string
	CompanyID   int
	ContentType string
	Content     []byte
	Name        string
}

var ContentType = struct {
	PDF string
}{
	PDF: "application/pdf",
}

/*
Upload a file to Firebase Storage

The file is uploaded to companyID/bucket/filename.ext

Return file URL.
*/
func UploadFile(file FileProps) (uri *string, url *string, err error) {
	ctx := context.Background()
	_uri := strconv.Itoa(file.CompanyID) + "/" + file.Bucket + "/" + file.Name
	uri = &_uri
	obj := inits.FirebaseStorage.Object(*uri)
	writer := obj.NewWriter(ctx)
	writer.ObjectAttrs.Metadata = map[string]string{"ContentType": file.ContentType}

	if _, err := io.Copy(writer, bytes.NewReader(file.Content)); err != nil {
		log.Default().Println("43 upload error:", err)
		return nil, nil, err
	}

	writer.Close()

	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Default().Println("50 upload error:", err)
		return nil, nil, err
	}

	if attrs, err := obj.Attrs(ctx); err != nil {
		log.Default().Println("55 upload error:", err)
		return nil, nil, err
	} else {
		return uri, &attrs.MediaLink, nil
	}
}


// Delete file from Firebase storage
func DeleteFile(uri string) error {
	ctx := context.Background()
	return inits.FirebaseStorage.Object(uri).Delete(ctx)
}