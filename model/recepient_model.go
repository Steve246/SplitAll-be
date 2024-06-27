package model

import "gorm.io/gorm"

type UserRecepient struct {
	gorm.Model
	MenuName   string `json:"menuName"`
	MenuPrice  string `json:"menuPrice"`
	AssignTo   string `json:"assignTo"`
	BankType   string `json:"bankType"`
	BankNumber string `json:"bankNumber"`
}
