package utils

import "testing"

func TestGetPhoneAndCountryCode(t *testing.T) {
	phone := "+62813222330"
	phoneNo, code := GetPhoneAndCountryCode(phone)

	expectedPhoneNo := "813222330"
	expectedCountryCode := "+62"
	if phoneNo != expectedPhoneNo {
		t.Fatalf("Expected phone %s, got %s", expectedPhoneNo, phone)
	}

	if code != expectedCountryCode {
		t.Fatalf("Expected phone %s, got %s", expectedCountryCode, code)
	}
}

func TestAuthenticatePassword(t *testing.T) {
	dummyPwd := "dolor1psum@"

	hashedPwd, err := HashPassword(dummyPwd)
	if err != nil {
		t.FailNow()
	}

	valid := ComparePasswordHash(dummyPwd, hashedPwd)
	if !valid {
		t.Fatal("password should be valid")
	}
}
