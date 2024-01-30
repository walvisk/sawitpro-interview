// This file contains types that are used in the repository layer.
package repository

import (
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	CountryCode string    `json:"country_code"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type UserLog struct {
	ID      int64     `json:"id"`
	LoginAt time.Time `json:"deleted_at"`
}
