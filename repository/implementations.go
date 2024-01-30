package repository

import "context"

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

	err = tx.QueryRow(
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
