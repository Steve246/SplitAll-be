package repository

import (
	"SplitAll/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type OcrReaderRepository interface {
	PostOcrData(imageData []byte, contentType string) (string, error)
}

type ocrReaderRepository struct {
	db        *gorm.DB
	ocrConfig config.OcrConfig
	// baseURL string
}

func (repo *ocrReaderRepository) PostOcrData(imageData []byte, contentType string) (string, error) {
	// Create the full URL

	fmt.Println("ini masuk PostOCrData")

	url := repo.ocrConfig.ApiUrl + repo.ocrConfig.ApiEndpoint

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(imageData))

	fmt.Println("ini error PostOCrData --> ", err)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/json")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func NewOcrReaderRepository(db *gorm.DB, ocrConfig config.OcrConfig) OcrReaderRepository {
	repo := new(ocrReaderRepository)
	repo.db = db
	repo.ocrConfig = ocrConfig
	return repo
}
