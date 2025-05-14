package helper

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var srv *drive.Service
var DriveService *drive.Service

func InitDriveService() error {
	credsB64 := os.Getenv("GOOGLE_CREDENTIALS_JSON_BASE64")
	if credsB64 == "" {
		return fmt.Errorf("GOOGLE_CREDENTIALS_JSON_BASE64 is not set")
	}

	credsBytes, err := base64.StdEncoding.DecodeString(credsB64)
	if err != nil {
		return fmt.Errorf("failed to decode base64 credentials: %v", err)
	}

	ctx := context.Background()
	srv, err = drive.NewService(ctx, option.WithCredentialsJSON(credsBytes))
	if err != nil {
		return fmt.Errorf("failed to initialize Google Drive: %v", err)
	}

	DriveService = srv
	return nil
}

// Upload file ke Google Drive dengan validasi null pointer
func UploadToDrive(file multipart.File, header *multipart.FileHeader, folderID string) (string, string, error) {
	if file == nil || header == nil {
		return "", "", errors.New("file atau header tidak boleh nil")
	}

	ext := filepath.Ext(header.Filename)
	filename := uuid.New().String() + ext

	f := &drive.File{
		Name:     filename,
		Parents:  []string{folderID},
		MimeType: header.Header.Get("Content-Type"),
	}

	res, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		return "", "", fmt.Errorf("failed to upload to drive: %v", err)
	}

	err = makePublic(res.Id)
	if err != nil {
		return "", "", fmt.Errorf("failed to make file public: %v", err)
	}

	return res.Id, getPublicURL(res.Id), nil
}

// Hapus file dari Google Drive
func DeleteFromDrive(fileID string) error {
	if fileID == "" {
		return nil
	}
	return srv.Files.Delete(fileID).Do()
}

// Buat URL publik dari fileID
func getPublicURL(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/thumbnail?id=%s", fileID)
}

// Buat file publik
func makePublic(fileID string) error {
	perm := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err := srv.Permissions.Create(fileID, perm).Do()
	return err
}

// Ambil URL publik dari ID
func PublicImageURLDrive(fileID string) string {
	if fileID == "" {
		return ""
	}
	return getPublicURL(fileID)
}
