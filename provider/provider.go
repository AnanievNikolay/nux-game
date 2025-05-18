package provider

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/common/db"
	"github.com/AnanievNikolay/nux-game/common/lifecycle"
	"github.com/AnanievNikolay/nux-game/delivery/http"
)

type Provider struct {
	err error

	ctx    context.Context
	cfg    *config.Config
	logger *log.Entry

	lifecycleHub *lifecycle.Hub

	container *dig.Container
}

func NewProvider(
	ctx context.Context,
	cfg *config.Config,
	logger *log.Entry,
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

	p.provide(func() *log.Entry {
		return p.logger
	})

	// sqLite
	p.provide(func(
		logger *log.Entry,
		cfg *config.Config,
	) (*db.Connector, error) {
		return db.NewSQLiteConnector(
			logger,
			cfg,
		)
	}, dig.As(
		new(db.SQLiteDB),
	))

	p.provide(http.NewDelivery)

	p.hooks()

	return p.container, p.err
}
