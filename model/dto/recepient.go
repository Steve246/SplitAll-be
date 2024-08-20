package dto

type RecepientResponse struct {
	AssignPerson string
	MenuDetail   []RecepientDetail
	BankType     string
	BankNumber   string
}

type RecepientAssign struct {
	AssignPerson string
	BankType     string
	BankNumber   string
}

type RecepientDetail struct {
	MenuName  string `json:"menuName"`
	MenuPrice string `json:"menuPrice"`
}
