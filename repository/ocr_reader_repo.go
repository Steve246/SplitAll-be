package repository

import (
	"SplitAll/config"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

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
	fmt.Println("OCR API URL:", url)

	// Open the file and log potential errors
	file, err := imageData.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a buffer to hold the multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the form file part explicitly and set content type manually
	part, err := writer.CreatePart(
		textproto.MIMEHeader{
			"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, imageData.Filename)},
			"Content-Type":        []string{"image/jpeg"}, // Explicitly set the Content-Type for the file
		})
	if err != nil {
		return "", fmt.Errorf("failed to create form file part: %w", err)
	}

	// Copy the file content to the form file
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file data: %w", err)
	}

	// Close the writer to finalize the multipart form data
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create the POST request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set the correct Content-Type header for multipart form data
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	// Log the actual Content-Type being sent
	fmt.Println("Sending Content-Type:", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer response.Body.Close()

	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(bodyResponse), nil
}

func NewOcrReaderRepository(db *gorm.DB, ocrConfig config.OcrConfig) OcrReaderRepository {
	repo := new(ocrReaderRepository)
	repo.db = db
	repo.ocrConfig = ocrConfig
	return repo
}
