package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type JwtPostgres struct {
	db *sqlx.DB
}

func NewJwtPostgres(db *sqlx.DB) *JwtPostgres {
	return &JwtPostgres{db: db}
}

func (r *JwtPostgres) CreateNewRefreshToken(userId int, newRefreshToken string) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", refreshTokensTable),
		userId,
	)
	if err != nil {
		return "", fmt.Errorf("failed to delete old refresh token: %w", err)
	}

	var addedRefreshToken string
	err = tx.QueryRow(
		fmt.Sprintf("INSERT INTO %s (refresh_token, user_id) VALUES ($1, $2) RETURNING refresh_token", refreshTokensTable),
		newRefreshToken, userId,
	).Scan(&addedRefreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to insert new refresh token: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return addedRefreshToken, nil
}

func (r *JwtPostgres) RemoveRefreshToken(tokenString string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token=$1", refreshTokensTable)
	_, err := r.db.Exec(query, tokenString)
	if err != nil {
		return err
	}

	return nil
}
