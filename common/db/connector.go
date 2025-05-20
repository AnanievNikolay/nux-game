package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Connector struct {
	logger *logrus.Entry
	dbType string
	db     *sqlx.DB
}
