package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/domain"
)

type UserRepository interface {
	SaveUserTX(
		ctx context.Context,
		tx *sqlx.Tx,
		user *domain.User,
	) error
}

type TokenRepositopry interface {
	SaveTokenTX(ctx context.Context, tx *sqlx.Tx, token *domain.Token) error
}

type Connector interface {
	GetDB(ctx context.Context) *sqlx.DB
}

type UnitOfWork struct {
	con Connector

	userRepository   UserRepository
	tokenRepositopry TokenRepositopry

	ttl int
}

func NewUnitOwWork(
	con Connector,

	cfg *config.Config,

	userRepository UserRepository,
	tokenRepositopry TokenRepositopry,
) *UnitOfWork {
	return &UnitOfWork{
		con: con,

		ttl: cfg.DB.SQLite.TTL,

		userRepository:   userRepository,
		tokenRepositopry: tokenRepositopry,
	}
}
