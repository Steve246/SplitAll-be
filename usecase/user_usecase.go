package usecase

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/repository"
	"SplitAll/utils"
	"fmt"
	"mime/multipart"
)

type UserUsecase interface {
	GetOcrInfo(file *multipart.FileHeader) (string, error)
	UserSendRecepeint(images []model.UserRecepient) ([]dto.RecepientResponse, error)
	SaveImageURL(file *multipart.FileHeader) (string, error)
}

type userUsecase struct {
	ocrRepo       repository.OcrReaderRepository
	imageRepo     repository.ImageUploadRepository
	recepientRepo repository.RecepientRepository
}

func (u *userUsecase) GetOcrInfo(file *multipart.FileHeader) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("function panicked: %v", r)
		}
	}()

	// Determine the content type
	contentType := file.Header.Get("Content-Type")

	// Call the repository method to handle the OCR request
	ocrResult, err := u.ocrRepo.PostOcrData(file, contentType)
	if err != nil {
		return "", fmt.Errorf("failed to post OCR data: %w", err)
	}

	return ocrResult, nil
}

func (u *userUsecase) SaveImageURL(file *multipart.FileHeader) (string, error) {

	imageUrlNew, err := u.imageRepo.Create(file)

	if err != nil {
		return "", utils.ImageTypeError()
	}

	return imageUrlNew, nil
}

func (u *userUsecase) UserSendRecepeint(images []model.UserRecepient) ([]dto.RecepientResponse, error) {
	resultConvertText, err := u.recepientRepo.ConvertText(images)

	fmt.Println("ini hasil resultConvertText --> ", resultConvertText)

	if err != nil {
		return []dto.RecepientResponse{}, nil
	}

	return resultConvertText, nil

}

func NewUserUsecase(imageRepo repository.ImageUploadRepository, recepientRepo repository.RecepientRepository, ocrRepo repository.OcrReaderRepository) UserUsecase {
	usecase := new(userUsecase)
	usecase.imageRepo = imageRepo
	usecase.recepientRepo = recepientRepo
	usecase.ocrRepo = ocrRepo

	return usecase
}
