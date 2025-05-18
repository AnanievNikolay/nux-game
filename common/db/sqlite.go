package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/common/lifecycle"
)

const (
	dialect = `sqlite3`
)

type SQLiteDB interface {
	lifecycle.Service
}

func NewSQLiteConnector(
	logger *log.Entry,
	cfg *config.Config,
) (*Connector, error) {
	c := &Connector{
		logger: logger.WithFields(log.Fields{
			"layer":     "connector",
			"connector": dialect,
		}),
		dbType: dialect,
	}

	var err error

	if c.db, err = GetDBInstance(cfg.DB.SQLite); err != nil {
		c.logger.Fatalf("error while init connection to %s db: %v", c.dbType, err)
		return nil, err
	}

	c.logger.Infof("Connection to %s db is ok", c.dbType)

	return c, nil
}

func GetDBInstance(config *config.SQLite) (sqlxDB *sqlx.DB, err error) {
	if err := os.MkdirAll(config.FileFolder, os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	sqlxDB, err = sqlx.Open(dialect, config.FileFolder+"/"+config.FileName)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", err)
	}

	if err = migrateDB(config, sqlxDB); err != nil {
		return nil, fmt.Errorf("migrateDB: %w", err)
	}

	return
}

func migrateDB(config *config.SQLite, db *sqlx.DB) error {
	if config.MigratePath == "" {
		return nil
	}

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}
	return goose.Up(db.DB, config.MigratePath)
}

func (c *Connector) Stop(_ context.Context) error {
	if err := c.db.Close(); err != nil {
		c.logger.Errorf("error while close connection to %s DB: %s", c.dbType, err)
	}
	c.logger.Infof("Stopped %s connector", c.dbType)
	return nil
}

func (c *Connector) GetDB(context.Context) *sqlx.DB {
	return c.db
}
