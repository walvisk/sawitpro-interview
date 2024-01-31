package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
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

func (r *Repository) FindUserByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, full_name, password,  phone, country_code, created_at, updated_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	var user User
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
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

func (r *Repository) UpdateUser(ctx context.Context, u *User, fullName, phone string) error {
	var setClause strings.Builder
	var setParams []interface{}

	if fullName != "" {
		setClause.WriteString("full_name = $2, ")
		setParams = append(setParams, u.FullName)
	}

	if phone != "" {
		setClause.WriteString("phone = $3, ")
		setParams = append(setParams, u.Phone)
	}

	setClause.WriteString("updated_at = $4")
	setParams = append(setParams, time.Now())

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $1
		RETURNING id, full_name, phone, country_code, created_at, updated_at
	`, setClause.String())
	params := append([]interface{}{u.ID}, setParams...)

	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, query, u.ID, params).Scan(
		&u.ID,
		&u.FullName,
		&u.Phone,
		&u.CountryCode,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateUserLog(ctx context.Context, u *User) error {
	query := `
		INSERT INTO user_logs (user_id, login_at)
		VALUES ($1, $2)
	`

	_, err := r.Db.ExecContext(ctx, query, u.ID, time.Now())
	if err != nil {
		return err
	}

	return nil
}
