package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

func LogTimeSpent(
	logger *logrus.Entry,
	method string,
) func() {
	start := time.Now()
	logger.
		WithField("checkpoint", method).
		WithField("isCheckpoint", true).
		WithField("start", time.Now()).
		Infof("%s() started", method)
	return func() {
		took := time.Since(start)
		logger.
			WithField("checkpoint", method).
			WithField("isCheckpoint", true).
			WithField("finish", time.Now()).
			Infof("%s() finished; took = %s", method, took)
	}
}
