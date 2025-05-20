package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/token/sqlite/dto"
)

func (r *Repository) SaveTokenTX(ctx context.Context, tx *sqlx.Tx, token *domain.Token) error {
	q := `
	INSERT INTO user_token (
		user_id, token, expires_at
	) VALUES (
		:user_id, :token, :expires_at
	);`

	dtoToken := dto.NewToken(token)

	if _, err := tx.NamedExecContext(
		ctx,
		q,
		dtoToken,
	); err != nil {
		return fmt.Errorf("tx.NamedExecContext: %w", err)
	}

	return nil
}
