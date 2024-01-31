package handler

import (
	"testing"
)

func TestUserValidator(t *testing.T) {
	tests := []struct {
		name      string
		validator UserValidator
		assert    func(bool)
	}{
		{
			name: "when given correct params, HasError returns false",
			validator: UserValidator{
				FullName: "Dolor Ipsume Sit Amet",
				Phone:    "+628132456889",
				Password: "d0lorIpsum@",
			},
			assert: func(b bool) {
				if b == true {
					t.Fatalf("param should be valid")
				}
			},
		},
		{
			name: "when given wrong Phone, HasError returns true",
			validator: UserValidator{
				FullName: "Dolor Ipsume Sit Amet",
				Phone:    "+618132456889",
				Password: "d0lorIpsum@",
			},
			assert: func(b bool) {
				if b == false {
					t.Fatalf("param should not be valid")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.validator.ValidateFullName().ValidatePhone().ValidatePassword()

			err := tt.validator.HasError()
			tt.assert(err)
		})
	}
}
