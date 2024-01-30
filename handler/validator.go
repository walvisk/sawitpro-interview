package handler

import (
	"errors"
	"regexp"

	"github.com/SawitProRecruitment/UserService/generated"
)

type UserValidator struct {
	FullName  string
	Phone     string
	Password  string
	UserError *generated.ErrorResponse
}

func (us *UserValidator) HasError() bool {
	if len(us.UserError.Fields) > 0 {
		return true
	}

	return false
}

func (us *UserValidator) Validate() {
	errDetails := make([]*generated.ErrorResponseDetail, 0)
	if err := ValidateFullName(us.FullName); err != nil {
		errDetails = append(errDetails, &generated.ErrorResponseDetail{
			Error: err.Error(),
			Field: "full_name",
		})
	}

	if err := ValidatePhone(us.Phone); err != nil {
		errDetails = append(errDetails, &generated.ErrorResponseDetail{
			Error: err.Error(),
			Field: "phone",
		})
	}

	if err := ValidatePassword(us.Password); err != nil {
		errDetails = append(errDetails, &generated.ErrorResponseDetail{
			Error: err.Error(),
			Field: "password",
		})
	}

	var errResponse *generated.ErrorResponse
	if len(errDetails) > 0 {
		errResponse.Kind = "BadRequest"
		errResponse.Message = "Invalid Request Format"
		errResponse.Fields = errDetails
	}

	us.UserError = errResponse
}

func ValidateFullName(fullName string) error {
	if fullName != "" && (len(fullName) < 3 || len(fullName) > 60) {
		return errors.New("must be between 3 and 60 characters")
	}
	return nil
}

func ValidatePhone(phone string) error {
	if len(phone) < 3 {
		return errors.New("phone number must be between 10 and 13 characters")
	}

	countryCode, phoneNumber := phone[:3], phone[3:]
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 || countryCode != "+62" {
		return errors.New("phone number must be between 10 and 13 characters and start with '+62'")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 || len(password) > 64 ||
		!containsCapital(password) ||
		!containsNumber(password) ||
		!containsSpecial(password) {
		return errors.New("password must be between 6 and 64 characters and contain at least 1 capital letter, 1 number, and 1 special character")
	}
	return nil
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
