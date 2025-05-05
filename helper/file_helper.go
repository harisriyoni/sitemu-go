package helper

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

func SaveUploadedFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(header.Filename)

	// Gunakan UUID untuk nama file
	uniqueName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	publicFolder := filepath.Join("public", folder)
	fullPath := filepath.Join(publicFolder, uniqueName)

	err := os.MkdirAll(publicFolder, os.ModePerm)
	if err != nil {
		return "", err
	}

	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	_, err = out.ReadFrom(file)
	if err != nil {
		return "", err
	}

	return uniqueName, nil
}

func DeleteFile(folder, filename string) error {
	if filename == "" {
		return nil
	}
	fullPath := filepath.Join("public", folder, filename)
	return os.Remove(fullPath)
}

func PublicImageURL(folder, filename string) string {
	// Ganti "localhost:8080" jika kamu pakai domain/port berbeda
	return fmt.Sprintf("http://localhost:8080/public/%s/%s", folder, filename)
}

func GenerateRandomNumber() int {
	return rand.Intn(100000)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ReplaceUploadedFile(
	oldFilename string,
	newFile multipart.File,
	newHeader *multipart.FileHeader,
	folder string,
) (string, error) {
	// Hapus file lama
	err := DeleteFile(folder, oldFilename)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	// Simpan file baru
	newFilename, err := SaveUploadedFile(newFile, newHeader, folder)
	if err != nil {
		return "", err
	}
	return newFilename, nil
}

func CopyFile(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}
