package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func (c *Config) readConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// OCR Config

	ocrKeys := os.Getenv("OCR_API_KEYS")
	ocrUrl := os.Getenv("OCR_API_URL")
	c.OcrConfig = OcrConfig{ApiKeys: ocrKeys, ApiUrl: ocrUrl}

	// API Config start here
	// api := os.Getenv("API_URL")
	// c.ApiConfig = ApiConfig{Url: api}

	// DB Config start here
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	c.DbConfig = DbConfig{dsn}

	// c.FilePathConfig = FilePathConfig{FilePath: os.Getenv("FILE_PATH")}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}

func InitConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
