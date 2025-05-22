package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/uow/utils"
)

func (uow *UnitOfWork) UpdateToken(
	ctx context.Context,
	oldToken string,
	newToken *domain.Token,
) error {
	uowCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(uow.ttl)*time.Second)
	defer cancelFunc()

	tx, err := uow.con.GetDB(uowCtx).BeginTxx(uowCtx, nil)
	if err != nil {
		return fmt.Errorf("sqlx.BeginTxx: %w", err)
	}

	if err := uow.tokenRepositopry.DeactivateTX(uowCtx, tx, oldToken); err != nil {
		return utils.HandleTXError(tx, "tokenRepositopry.DeactivateTX", err)
	}

	if err := uow.tokenRepositopry.SaveTokenTX(uowCtx, tx, newToken); err != nil {
		return utils.HandleTXError(tx, "tokenRepositopry.SaveTokenTX", err)
	}

	return tx.Commit()
}
