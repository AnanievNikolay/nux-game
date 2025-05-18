package db

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Connector struct {
	logger *log.Entry
	dbType string
	db     *sqlx.DB
}
