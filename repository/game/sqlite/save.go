package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/game/sqlite/dto"
)

func (r *Repository) Save(ctx context.Context, game *domain.Game) (int, error) {
	gCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(r.ttl)*time.Second)
	defer cancelFunc()

	q := `INSERT INTO games (
			user_id, number, is_win, prize, created_at
		) VALUES (
			:user_id, :number, :is_win, :prize, :created_at
		);`

	dtoGame := dto.NewGame(game)

	res, err := r.con.GetDB(ctx).NamedExecContext(gCtx, q, dtoGame)
	if err != nil {
		return 0, fmt.Errorf("sqlx.NamedExecContext: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
