package repository

import (
	"SplitAll/model"
	"SplitAll/model/dto"
	"SplitAll/utils"

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

	for _, img := range images {
		menu := dto.RecepientDetail{
			MenuName:  img.MenuName,
			MenuPrice: img.MenuPrice,
		}
		menuGroups[img.AssignTo] = append(menuGroups[img.AssignTo], menu)
	}

	// Create the final response
	var result []dto.RecepientResponse
	for assignee, menus := range menuGroups {
		resp := dto.RecepientResponse{
			AssignPerson: assignee,
			MenuDetail:   menus,
			BankType:     result[0].BankType,   // Assuming the bank type is the same for all menus
			BankNumber:   result[0].BankNumber, // Assuming the bank number is the same for all menus
		}
		result = append(result, resp)
	}

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
