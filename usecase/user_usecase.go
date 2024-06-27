package usecase

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/repository"
	"SplitAll/utils"
	"mime/multipart"
)

type UserUsecase interface {
	UserSendRecepeint(images []model.UserRecepient) ([]dto.RecepientResponse, error)
	SaveImageURL(file *multipart.FileHeader) (string, error)
}

type userUsecase struct {
	imageRepo     repository.ImageUploadRepository
	recepientRepo repository.RecepientRepository
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
