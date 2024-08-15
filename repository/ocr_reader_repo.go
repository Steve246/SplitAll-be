package repository

import "gorm.io/gorm"

type OcrReaderRepository interface {
}

type ocrReaderRepository struct {
	db *gorm.DB
}

func NewOcrReaderRepository(db *gorm.DB) OcrReaderRepository {
	repo := new(ocrReaderRepository)
	repo.db = db
	return repo
}
