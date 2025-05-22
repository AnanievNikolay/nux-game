package sqlite

import (
	"context"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/jmoiron/sqlx"
)

type TokenRepositopry interface {
	SaveTokenTX(ctx context.Context, tx *sqlx.Tx, token *domain.Token) error
	DeactivateTX(ctx context.Context, tx *sqlx.Tx, token string) error
}

type Connector interface {
	GetDB(ctx context.Context) *sqlx.DB
}

type UnitOfWork struct {
	con Connector

	tokenRepositopry TokenRepositopry

	ttl int
}

func NewUnitOwWork(
	con Connector,

	cfg *config.Config,

	tokenRepositopry TokenRepositopry,
) *UnitOfWork {
	return &UnitOfWork{
		con: con,

		ttl: cfg.DB.SQLite.TTL,

		tokenRepositopry: tokenRepositopry,
	}
}
