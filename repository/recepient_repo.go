package repository

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/utils"
	"fmt"

	"gorm.io/gorm"
)

type RecepientRepository interface {
	ConvertText(images []model.UserRecepient) ([]dto.RecepientResponse, error)
}

type recepientRepository struct {
	db *gorm.DB
}

func (r *recepientRepository) ConvertText(images []model.UserRecepient) ([]dto.RecepientResponse, error) {
	// Create a map to group menus by assignee
	menuGroups := make(map[string][]dto.RecepientDetail)
	bankInfo := make(map[string]dto.RecepientResponse) // To store BankType and BankNumber for each assignee

	for _, img := range images {
		fmt.Println("ini img -->? ", img)
		menu := dto.RecepientDetail{
			MenuName:  img.MenuName,
			MenuPrice: img.MenuPrice,
		}

		// Add the menu to the correct assignee group
		menuGroups[img.AssignTo] = append(menuGroups[img.AssignTo], menu)

		// Store bank information (assuming it's the same for each assignee)
		if _, exists := bankInfo[img.AssignTo]; !exists {
			bankInfo[img.AssignTo] = dto.RecepientResponse{
				BankType:   img.BankType,
				BankNumber: img.BankNumber,
			}
		}
	}

	fmt.Println("ini masuk 2 -->", menuGroups)

	// Create the final response
	var result []dto.RecepientResponse
	for assignee, menus := range menuGroups {
		fmt.Println(assignee, menus)

		// Fetch the bank info for this assignee
		bank := bankInfo[assignee]

		resp := dto.RecepientResponse{
			AssignPerson: assignee,
			MenuDetail:   menus,
			BankType:     bank.BankType,   // Get the BankType from the stored info
			BankNumber:   bank.BankNumber, // Get the BankNumber from the stored info
		}
		result = append(result, resp)
	}

	fmt.Println("ini masuk 3 -->", result)

	if len(result) == 0 {
		return nil, utils.NoMenuProvidedError()
	}

	return result, nil
}

func NewRecepientRepository(db *gorm.DB) RecepientRepository {
	repo := new(recepientRepository)
	repo.db = db
	return repo
}
