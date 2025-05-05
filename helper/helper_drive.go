package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var srv *drive.Service

func InitDriveService() error {
	var err error
	srv, err = drive.NewService(context.Background(), option.WithCredentialsFile("uploadsitemu-3488eb80fb4e.json"))
	return err
}

func UploadToDrive(file multipart.File, header *multipart.FileHeader, folderID string) (string, string, error) {
	ext := filepath.Ext(header.Filename)
	filename := uuidStr() + ext

	f := &drive.File{
		Name:     filename,
		Parents:  []string{folderID},
		MimeType: header.Header.Get("Content-Type"),
	}

	res, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		return "", "", err
	}

	err = makePublic(res.Id)
	if err != nil {
		return "", "", err
	}

	return res.Id, getPublicURL(res.Id), nil
}

func DeleteFromDrive(fileID string) error {
	return srv.Files.Delete(fileID).Do()
}

func getPublicURL(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/uc?id=%s", fileID)
}

func makePublic(fileID string) error {
	perm := &drive.Permission{
		Type: "anyone", Role: "reader",
	}
	_, err := srv.Permissions.Create(fileID, perm).Do()
	return err
}

func uuidStr() string {
	return uuid.New().String()
}

func PublicImageURLDrive(fileID string) string {
	if fileID == "" {
		return ""
	}
	return fmt.Sprintf("https://drive.google.com/uc?id=%s", fileID)
}
