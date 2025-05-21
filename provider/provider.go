package provider

import (
	"context"

	"go.uber.org/dig"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/common/db"
	"github.com/AnanievNikolay/nux-game/common/lifecycle"
	"github.com/AnanievNikolay/nux-game/delivery/http"
	"github.com/sirupsen/logrus"

	userHandler "github.com/AnanievNikolay/nux-game/delivery/http/handler/user"

	tokenService "github.com/AnanievNikolay/nux-game/service/token"
	userService "github.com/AnanievNikolay/nux-game/service/user"

	tokenRepository "github.com/AnanievNikolay/nux-game/repository/token/sqlite"
	userRepository "github.com/AnanievNikolay/nux-game/repository/user/sqlite"

	userUnitOfWork "github.com/AnanievNikolay/nux-game/uow/user/sqlite"
)

type Provider struct {
	err error

	ctx    context.Context
	cfg    *config.Config
	logger *logrus.Entry

	lifecycleHub *lifecycle.Hub

	container *dig.Container
}

func NewProvider(
	ctx context.Context,
	cfg *config.Config,
	logger *logrus.Entry,
) *Provider {
	return &Provider{
		container: dig.New(),

		ctx:    ctx,
		cfg:    cfg,
		logger: logger,

		lifecycleHub: lifecycle.NewHub(logger),
	}
}

func (p *Provider) Provide() (*dig.Container, error) {
	p.provide(func() *config.Config {
		return p.cfg
	})

	p.provide(func() context.Context {
		return p.ctx
	})

	p.provide(func() *logrus.Entry {
		return p.logger
	})

	// sqLite
	p.provide(func(
		logger *logrus.Entry,
		cfg *config.Config,
	) (*db.Connector, error) {
		return db.NewSQLiteConnector(
			logger,
			cfg,
		)
	}, dig.As(
		new(userRepository.Connector),
		new(tokenRepository.Connector),
		new(userUnitOfWork.Connector),

		new(db.SQLiteDB),
	))

	// services
	p.provide(userService.NewService, dig.As(
		new(userHandler.Service),
	))

	p.provide(tokenService.NewService, dig.As(
		new(userService.TokenService),
	))

	// unit of works
	p.provide(userUnitOfWork.NewUnitOwWork, dig.As(
		new(userService.UnitOfWork),
	))

	// repositories
	p.provide(userRepository.NewRepository, dig.As(
		new(userUnitOfWork.UserRepository),
		new(userService.Repository),
	))

	p.provide(tokenRepository.NewRepository, dig.As(
		new(userUnitOfWork.TokenRepositopry),
		new(tokenService.Repository),
	))

	p.provide(http.NewDelivery)

	// handlers
	p.provide(userHandler.NewHandler)

	p.hooks()

	return p.container, p.err
}
