package dto

import (
	"unicode"
)

type (
	RegisterRequest struct {
		Name        string `json:"name" form:"name"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
		Password    string `json:"password" form:"password"`
	}

	RegisterResponse struct {
		ID int `json:"id"`
	}

	RegisterValidationErrorResponse struct {
		Name        string `json:"name,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Password    string `json:"password,omitempty"`
	}

	LoginRequest struct {
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
		Password    string `json:"password" form:"password" binding:"required"`
	}

	LoginResponse struct {
		ID    int    `json:"id"`
		Token string `json:"token"`
	}

	ProfileResponse struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}

	ProfileUpdateRequest struct {
		Name        string `json:"name" form:"name"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
	}
)

func (r RegisterRequest) Validate() RegisterValidationErrorResponse {
	var validationErrResponse RegisterValidationErrorResponse

	validationErrResponse.Name = validateName(r.Name)
	validationErrResponse.PhoneNumber = validatePhoneNumber(r.PhoneNumber)
	validationErrResponse.Password = validatePassword(r.Password)

	return validationErrResponse
}

func (r ProfileUpdateRequest) Validate() RegisterValidationErrorResponse {
	var validationErrResponse RegisterValidationErrorResponse

	validationErrResponse.Name = validateName(r.Name)
	validationErrResponse.PhoneNumber = validatePhoneNumber(r.PhoneNumber)

	return validationErrResponse
}

func validateName(name string) string {
	if name == "" {
		return "name must not be empty"
	} else if len(name) < 3 || len(name) > 60 {
		return "name must be at minimum 3 characters and maximum 60 characters"
	}

	return ""
}

func validatePhoneNumber(phone_number string) string {
	if phone_number == "" {
		return "phone number must not be empty"
	} else if len(phone_number) < 10 || len(phone_number) > 13 {
		return "phone number must be at minimum 10 characters and maximum 13 characters"
	} else if phone_number[0:3] != "+62" {
		return "phone number must start with the Indonesia country code '+62'"
	}

	return ""
}

func validatePassword(password string) string {
	if password == "" {
		return "password must not be empty"
	} else if len(password) < 6 || len(password) > 64 {
		return "password must be at minimum 6 characters and maximum 64 characters"
	} else if !isCapital(password) || !isDigit(password) || !isSpecialChar(password) {
		return "passwords must be containing at least 1 capital character, 1 number, and 1 special character"
	}

	return ""
}

func isCapital(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func isSpecialChar(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
