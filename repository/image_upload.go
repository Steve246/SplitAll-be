package repository

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageUploadRepository interface {
	Create(file *multipart.FileHeader) (string, error)
}

type imageUploadRepository struct {
	db *gorm.DB
}

func (r *imageUploadRepository) Create(file *multipart.FileHeader) (string, error) {
	// Validate the file format (JPEG)

	if !isValidImageFile(file) {
		return "", errors.New("invalid image format")
	}

	// Generate a unique filename (e.g., UUID)
	filename := generateUniqueFilename(file.Filename)

	// Save the file to the "file" folder
	filePath := filepath.Join("file", filename)
	if err := saveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	// Insert the image URL into the database
	// query := "INSERT INTO imagesuser (image_url) VALUES (?)"
	// if err := r.db.Exec(query, file.Filename).Error; err != nil {
	// 	// Rollback: Delete the saved file if database insertion fails
	// 	os.Remove(filePath)
	// 	return "", err
	// }

	return filePath, nil
}

func isValidImageFile(file *multipart.FileHeader) bool {
	ext := filepath.Ext(file.Filename)
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".jpeg"
}

func generateUniqueFilename(originalFilename string) string {
	// Generate a UUID (Universally Unique Identifier)
	uniqueID := uuid.New()

	// Get the file extension (e.g., ".jpg" or ".jpeg")
	ext := filepath.Ext(originalFilename)
	ext = strings.ToLower(ext)

	// Combine the UUID and extension to create the unique filename
	return uniqueID.String() + ext
}

func saveUploadedFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}

	fmt.Println("ini nama file --> ", dst)

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	os.Remove(dst.Name())

	return nil
}

func NewImageUploadRepository(db *gorm.DB) ImageUploadRepository {
	repo := new(imageUploadRepository)
	repo.db = db
	return repo
}
