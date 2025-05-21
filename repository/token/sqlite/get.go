package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/token/sqlite/dto"
)

func (r *Repository) GetToken(ctx context.Context, token string) (*domain.Token, error) {
	rCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(r.ttl)*time.Second)
	defer cancelFunc()

	q := `SELECT user_id, token, expires_at
			FROM user_token
		WHERE token = ?;`

	var dtoToken dto.Token

	if err := r.con.GetDB(rCtx).GetContext(rCtx, &dtoToken, q, token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("sqlx.GetContext: %w", err)
	}

	return dtoToken.ToDomain(), nil
}
