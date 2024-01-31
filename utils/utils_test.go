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

	tests := []struct {
		name      string
		pwd       string
		assertion func(bool)
	}{
		{
			name: "when given correct password, return true",
			pwd:  dummyPwd,
			assertion: func(b bool) {
				if !b {
					t.Fatal("password should be valid")
				}
			},
		},
		{
			name: "when given wrong password, return false",
			pwd:  "dummyPwd",
			assertion: func(b bool) {
				if b {
					t.Fatal("password should be invalid")
				}
			},
		},
	}

	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			valid := ComparePasswordHash(tC.pwd, hashedPwd)
			tC.assertion(valid)
		})
	}
}
