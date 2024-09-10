package usecase

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/repository"
	"SplitAll/utils"
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"
)

type UserUsecase interface {
	GetOcrInfo(file *multipart.FileHeader) (item []dto.Item, err error)
	UserSendRecepeint(images []model.UserRecepient) ([]dto.RecepientResponse, error)
	SaveImageURL(file *multipart.FileHeader) (string, error)
}

type userUsecase struct {
	ocrRepo       repository.OcrReaderRepository
	imageRepo     repository.ImageUploadRepository
	recepientRepo repository.RecepientRepository
}

func (u *userUsecase) GetOcrInfo(file *multipart.FileHeader) (item []dto.Item, err error) {
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
		return []dto.Item{}, fmt.Errorf("failed to post OCR data: %w", err)
	}

	// fmt.Println("ini hasil OCR --> ", ocrResult)

	// Regex pattern to capture item number, name, quantity, price, and amount
	pattern := `(\d+)\s+([A-Z\s]+)\s+(\d+)\s+(\d+(?:\.\d+)?)\s+(\d+(?:\.\d+)?)`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(ocrResult, -1)

	// Slice to hold the extracted items
	var items []dto.Item

	// Iterate over matches and store in struct
	for _, match := range matches {
		itemName := strings.TrimSpace(match[2])
		quantity := match[3]
		price := match[4]

		// Convert quantity and price to appropriate types
		var qty int
		var priceFloat float64
		fmt.Sscanf(quantity, "%d", &qty)
		fmt.Sscanf(price, "%f", &priceFloat)

		// Append the extracted item to the slice
		item := dto.Item{
			Name:     itemName,
			Quantity: qty,
			Price:    priceFloat,
		}
		items = append(items, item)
	}

	// Print the extracted items
	for _, item := range items {
		fmt.Printf("Item: %s, Quantity: %d, Price: %.2f\n", item.Name, item.Quantity, item.Price)
	}

	return items, nil
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
