package repository

import (
	"SplitAll/config"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"gorm.io/gorm"
)

type OcrReaderRepository interface {
	// PostOcrData(imageData []byte, contentType string) (string, error)
	PostOcrData(imageData *multipart.FileHeader, contentType string) (string, error)
}

type ocrReaderRepository struct {
	db        *gorm.DB
	ocrConfig config.OcrConfig
	// baseURL string
}

func (repo *ocrReaderRepository) PostOcrData(imageData *multipart.FileHeader, contentType string) (string, error) {
	url := repo.ocrConfig.ApiUrl + repo.ocrConfig.ApiEndpoint

	fmt.Println("ini url --> ", url)

	// Log the key and filename
	fmt.Printf("Using key: 'file', filename: %s\n", imageData.Filename)

	// Create a new buffer to hold the multipart data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Open the file to read its contents
	file, err := imageData.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a form-data field "file" and add the image data with the original filename
	part, err := writer.CreateFormFile("file", imageData.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}

	// Write the file data to the "file" field and calculate the size while copying
	size, err := io.Copy(part, file)
	if err != nil || size == 0 {
		return "", fmt.Errorf("failed to write image data or file is empty: %w", err)
	}

	// Log the number of bytes written to the form
	fmt.Printf("Bytes written to form: %d\n", size)

	// Close the writer to finalize the multipart form
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Print the buffer content length for debugging
	fmt.Printf("Buffer content length: %d\n", buf.Len())

	// Create a new POST request with the multipart form data
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set the correct content type for the multipart form data
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	fmt.Printf("Request Headers: %v\n", req.Header)
	// fmt.Printf("Request Body: %s\n", buf.String()) // Be careful with large files

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer response.Body.Close()

	// Print response status for debugging
	fmt.Printf("Response status: %s\n", response.Status)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Print response body for debugging
	fmt.Printf("Response body: %s\n", string(body))

	return string(body), nil
}

func NewOcrReaderRepository(db *gorm.DB, ocrConfig config.OcrConfig) OcrReaderRepository {
	repo := new(ocrReaderRepository)
	repo.db = db
	repo.ocrConfig = ocrConfig
	return repo
}
