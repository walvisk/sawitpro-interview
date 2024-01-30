package handler

import (
	"regexp"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/utils"
)

type UserValidator struct {
	FullName string
	Phone    string
	Password string
	Errors   []*generated.ErrorResponseDetail
}

func (uv *UserValidator) ValidateFullName() *UserValidator {
	if uv.FullName != "" && (len(uv.FullName) < 3 || len(uv.FullName) > 60) {
		uv.Errors = append(uv.Errors, &generated.ErrorResponseDetail{
			Error: "must be between 3 and 60 characters",
			Field: "full_name",
		})
	}
	return uv
}

func (uv *UserValidator) ValidatePhone() *UserValidator {
	if len(uv.Phone) < 3 {
		uv.Errors = append(uv.Errors, &generated.ErrorResponseDetail{
			Error: "phone number must be between 10 and 13 characters",
			Field: "phone",
		})

		return uv
	}

	countryCode, phoneNumber := utils.GetPhoneAndCountryCode(uv.Phone)
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 || countryCode != "+62" {
		uv.Errors = append(uv.Errors, &generated.ErrorResponseDetail{
			Error: "phone number must be between 10 and 13 characters and start with '+62'",
			Field: "phone",
		})
	}

	return uv
}

func (uv *UserValidator) ValidatePassword() *UserValidator {
	if len(uv.Password) < 6 || len(uv.Password) > 64 ||
		!containsCapital(uv.Password) ||
		!containsNumber(uv.Password) ||
		!containsSpecial(uv.Password) {
		uv.Errors = append(uv.Errors, &generated.ErrorResponseDetail{
			Error: "password must be between 6 and 64 characters and contain at least 1 capital letter, 1 number, and 1 special character",
			Field: "phone",
		})
	}

	return uv
}

func (uv *UserValidator) HasError() bool {
	if len(uv.Errors) > 0 {
		return true
	}

	return false
}

// containsCapital checks if the password contains at least 1 capital letter.
func containsCapital(s string) bool {
	re := regexp.MustCompile(`[A-Z]`)
	return re.MatchString(s)
}

// containsNumber checks if the password contains at least 1 number.
func containsNumber(s string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(s)
}

// containsSpecial checks if the password contains at least 1 special character.
func containsSpecial(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.MatchString(s)
}
