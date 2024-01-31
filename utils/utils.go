package utils

import "golang.org/x/crypto/bcrypt"

func GetPhoneAndCountryCode(phoneString string) (phone, code string) {
	return phoneString[3:], phoneString[:3]
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
