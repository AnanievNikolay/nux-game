package http

import (
	"fmt"
	"io"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/config"
)

type Delivery struct {
	logger *log.Entry
	e      *echo.Echo

	addr string
}

func NewDelivery(
	logger *log.Entry,

	cfg *config.Config,
) *Delivery {
	d := Delivery{
		logger: logger.WithFields(log.Fields{
			"layer":    "delivery",
			"delivery": "http",
		}),
		e: echo.New(),

		addr: fmt.Sprintf("%s:%d", cfg.Delivery.HTTP.Host, cfg.Delivery.HTTP.Port),
	}

	d.e.Logger.SetOutput(io.Discard)

	d.initRouter()

	return &d
}
