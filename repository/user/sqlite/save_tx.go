package sqlite

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/user/sqlite/dto"
)

func (r *Repository) SaveUserTX(
	ctx context.Context,
	tx *sqlx.Tx,
	user *domain.User,
) error {
	q := `
	INSERT INTO users (
		id, username, phone
	) VALUES (
		:id, :username, :phone
	);`

	dtoUser := dto.NewUser(user)

	if _, err := tx.NamedExecContext(
		ctx,
		q,
		dtoUser,
	); err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint &&
				sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return fmt.Errorf("user already exists (unique constraint): %w", err)
			}
		}
		return fmt.Errorf("tx.NamedExecContext: %w", err)
	}

	return nil
}
