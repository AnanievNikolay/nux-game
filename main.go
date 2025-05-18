package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/delivery/http"
	"github.com/AnanievNikolay/nux-game/provider"
	"github.com/sirupsen/logrus"
)

const (
	configEnvVar = "APP_CONFIG"
)

var (
	stopTimeoutDuration = 20 * time.Second
)

// @securityDefinitions.apikey ApiKeyAuth
// @type apiKey
// @in query
// @name token
func main() {
	time.Local = time.UTC

	appCtx := context.TODO()

	configPath := os.Getenv(configEnvVar)
	if configPath == "" {
		panic("APP_CONFIG environment variable is not set")
	}

	serviceConfig, err := config.PrepareConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf(
			"cfg.PrepareConfig: %s",
			err,
		))
	}

	multiWriter := io.MultiWriter(os.Stderr, serviceConfig.Logger)

	loggerInstance := logrus.New()
	loggerInstance.SetOutput(multiWriter)
	loggerInstance.SetLevel(logrus.DebugLevel)
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})

	logger := loggerInstance.WithField("logger", "nux-game")

	p := provider.NewProvider(
		appCtx,
		serviceConfig,
		logger,
	)

	diContainer, err := p.Provide()
	if err != nil {
		logger.Fatalf("cant init provider: %s", err)
	}

	// wait for system signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(
		stopChan,
		syscall.SIGINT,  // for local Ctrl+C
		syscall.SIGTERM, // for remote stop container
	)

	// start HTTP server
	if err = diContainer.Invoke(func(delivery *http.Delivery) {
		go func() {
			if err = delivery.Start(); err != nil {
				logger.Fatal(err)
			}
		}()
	}); err != nil {
		logger.Fatalf("cant start http server: %s", err)
	}

	// wait for system signal
	sig := <-stopChan

	logger.Infof("exit signal captured: %v", sig)
	logger.Infof("stopping all services (%s)...", stopTimeoutDuration)

	stopCtx, stopCtxCancel := context.WithTimeout(appCtx, stopTimeoutDuration)
	defer stopCtxCancel()

	p.Stop(stopCtx)

	if stopCtx.Err() != nil {
		logger.Warn("stopCtx.Err:", stopCtx.Err())
	}
}
