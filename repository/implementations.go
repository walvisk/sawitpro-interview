package repository

import (
	"context"
	"database/sql"
)

func (r *Repository) CreateUser(ctx context.Context, u *User) (int64, error) {
	query := `
		INSERT INTO users (full_name, phone, country_code, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	tx, err := r.Db.Begin()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRowContext(
		ctx,
		query,
		u.FullName,
		u.Phone, u.CountryCode,
		u.Password,
		u.CreatedAt, u.UpdatedAt).Scan(&u.ID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (r *Repository) FindUserByPhoneAndCountryCode(ctx context.Context, phone, countryCode string) (*User, error) {
	query := `
		SELECT id, full_name, password, phone, country_code, created_at, updated_at
		FROM users
		WHERE phone = $1
		AND country_code = $2
		LIMIT 1
	`

	var user User
	err := r.Db.QueryRowContext(ctx, query, phone, countryCode).Scan(
		&user.ID,
		&user.FullName,
		&user.Password,
		&user.Phone,
		&user.CountryCode,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
