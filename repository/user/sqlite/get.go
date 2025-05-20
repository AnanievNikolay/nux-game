package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repository) GetIDByUsernameAndPhone(
	ctx context.Context,
	username, phone string,
) (string, error) {
	q := `SELECT id FROM users WHERE username = ? AND phone = ?`

	var id string

	if err := r.con.GetDB(ctx).QueryRowContext(ctx, q, username, phone).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}

		return "", fmt.Errorf("sqlx.QueryRowContext: %w", err)
	}

	return id, nil
}
