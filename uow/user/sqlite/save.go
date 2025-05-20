package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/uow/utils"
)

func (uow *UnitOfWork) Save(
	ctx context.Context,
	user *domain.User,
	token *domain.Token,
) error {
	uowCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(uow.ttl)*time.Second)
	defer cancelFunc()

	tx, err := uow.con.GetDB(uowCtx).BeginTxx(uowCtx, nil)
	if err != nil {
		return fmt.Errorf("sqlx.BeginTxx: %w", err)
	}

	if err := uow.userRepository.SaveUserTX(uowCtx, tx, user); err != nil {
		return utils.HandleTXError(tx, "userRepository.SaveUserTX", err)
	}

	if err := uow.tokenRepositopry.SaveTokenTX(uowCtx, tx, token); err != nil {
		return utils.HandleTXError(tx, "tokenRepositopry.SaveTokenTX", err)
	}

	return tx.Commit()
}
