package sqlite

import (
	"context"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/jmoiron/sqlx"
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
