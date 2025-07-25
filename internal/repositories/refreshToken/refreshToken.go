package refreshTokenRepository

import (
	"github.com/jmoiron/sqlx"
)

const refreshTokensTable = "refresh_tokens"

type RefreshToken interface {
	CreateToken(userId int, newRefreshToken string) (string, error)
	RemoveToken(tokenString string) error
}

type RefreshTokenRepository struct {
	db *sqlx.DB
}

func NewRefreshTokenRepository(db *sqlx.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}
