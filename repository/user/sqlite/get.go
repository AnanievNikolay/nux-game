package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/user/sqlite/dto"
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

func (r *Repository) GetByID(
	ctx context.Context,
	id string,
) (*domain.User, error) {
	q := `SELECT
				id,
				username,
				phone
			FROM
				users
			WHERE
				id = ?`

	var dtoUser dto.User

	if err := r.con.GetDB(ctx).GetContext(ctx, &dtoUser, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("sqlx.GetContext: %w", err)
	}

	return dtoUser.ToDomain(), nil
}
