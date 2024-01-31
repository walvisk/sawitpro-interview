// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	CreateUser(c context.Context, u *User) (int64, error)
	FindUserByPhoneAndCountryCode(c context.Context, phone, countryCode string) (*User, error)
	FindUserByID(c context.Context, id int64) (*User, error)
	UpdateUser(c context.Context, u *User, fullName, phone string) error
	CreateUserLog(c context.Context, u *User) error
}
