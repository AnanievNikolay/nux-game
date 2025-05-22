package sqlite

import (
	"context"
	"fmt"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/game/sqlite/dto"
)

func (r *Repository) GetHistoryByToken(ctx context.Context, token string) ([]domain.Game, error) {
	q := `SELECT
			g.id,
			g.user_id,
			g.is_win,
			g.number,
			g.prize,
			g.created_at
		FROM
			games AS g
		INNER JOIN user_token AS ut ON ut.user_id = g.user_id
			AND ut.token = ?
		ORDER BY
			g.created_at DESC
		LIMIT 3`

	var dest dto.Games

	if err := r.con.GetDB(ctx).SelectContext(ctx, &dest, q, token); err != nil {
		return nil, fmt.Errorf("sqlx.SelectContext: %w", err)
	}

	return dest.ToDomain(), nil
}
