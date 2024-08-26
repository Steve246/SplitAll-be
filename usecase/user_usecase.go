package usecase

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/repository"
	"SplitAll/utils"
	"fmt"
	"io/ioutil"
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

func (u *userUsecase) GetOcrInfo(file *multipart.FileHeader) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// Read the file content into a byte slice
	imageData, err := ioutil.ReadAll(src)

	if err != nil {
		return "", fmt.Errorf("failed to read file content: %w", err)
	}

	// Determine the content type
	contentType := file.Header.Get("Content-Type")

	fmt.Println("sampe sini masih jalan --> ", imageData)

	// Call the PostOcrData function
	ocrResult, errOcr := u.ocrRepo.PostOcrData(imageData, contentType)

	if errOcr != nil {
		return "", utils.ApiOcrError()
	}

	// Save the image URL or handle the OCR result as needed
	// imageUrlNew, err := u.imageRepo.Create(file)
	// if err != nil {
	// 	return "", utils.ImageTypeError()
	// }

	// Return the image URL and/or OCR result
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

func NewUserUsecase(imageRepo repository.ImageUploadRepository, recepientRepo repository.RecepientRepository) UserUsecase {
	usecase := new(userUsecase)
	usecase.imageRepo = imageRepo
	usecase.recepientRepo = recepientRepo

	return usecase
}
