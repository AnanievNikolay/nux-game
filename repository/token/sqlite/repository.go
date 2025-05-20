package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/AnanievNikolay/nux-game/common/config"
)

type Connector interface {
	GetDB(ctx context.Context) *sqlx.DB
}

type Repository struct {
	con Connector

	ttl int
}

func NewRepository(
	con Connector,

	cfg *config.Config,
) *Repository {
	return &Repository{
		con: con,

		ttl: cfg.DB.SQLite.TTL,
	}
}
