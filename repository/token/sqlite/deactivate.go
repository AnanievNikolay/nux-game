package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const deactivateQuery = `UPDATE user_token SET expires_at = datetime('now', '-1 second') WHERE token = ?;`

func (r *Repository) Deactivate(ctx context.Context, token string) error {
	if _, err := r.con.GetDB(ctx).ExecContext(ctx, deactivateQuery, token); err != nil {
		return fmt.Errorf("sqlx.ExecContext: %w", err)
	}

	return nil
}

func (r *Repository) DeactivateTX(ctx context.Context, tx *sqlx.Tx, token string) error {
	if _, err := tx.ExecContext(ctx, deactivateQuery, token); err != nil {
		return fmt.Errorf("sqlx.ExecContext: %w", err)
	}

	return nil
}
