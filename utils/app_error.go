package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

// api ocr error

func ApiOcrError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "API OCR Connection error",
		ErrorType:    http.StatusBadRequest,
	}
}

// recepient error

func NoMenuProvidedError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "No Menu Provided",
		ErrorType:    http.StatusBadRequest,
	}
}

// merchant crud

func InsertMerchantDataError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Merchant data failed to insert",
		ErrorType:    http.StatusBadRequest,
	}
}

// image upload

func UploadImageFileLimitation() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "File must be between 10KB and 2MB",
		ErrorType:    http.StatusBadRequest,
	}
}

func UploadImageTypeError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Image Type need to be .jpg or .jpeg",
		ErrorType:    http.StatusBadRequest,
	}
}

func UploadImageError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Upload Image Error",
		ErrorType:    http.StatusBadRequest,
	}
}

func ImageTypeError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Image type has to be .jpg or .jpeg",
		ErrorType:    http.StatusBadRequest,
	}
}

// login

func PasswordCannotBeEncodeError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password cannot be encode",
		ErrorType:    http.StatusBadRequest,
	}
}

func WrongInputPassword() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password Not Match",
		ErrorType:    http.StatusBadRequest,
	}
}

func UserDataNotFound() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "User Data Not Found",
		ErrorType:    http.StatusBadRequest,
	}
}

// Register

func ReqBodyNotValidError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Request body not valid",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrorDuplicatePassword() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "password already used",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrrorDuplicateEmail() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "email already used",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrorValidateEmail() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "not pass validation email",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrorValidatePassword() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "not pass validation password",
		ErrorType:    http.StatusBadRequest,
	}
}

func ErrorValidateName() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "not pass validation name",
		ErrorType:    http.StatusBadRequest,
	}
}

func CreateUserError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "user cannot be created",
		ErrorType:    http.StatusBadRequest,
	}
}
