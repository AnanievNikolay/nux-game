package http

import (
	"fmt"
	"io"

	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/delivery/http/handler/user"
)

type Delivery struct {
	logger *logrus.Entry

	e *echo.Echo

	addr string

	userHandler *user.Handler
}

func NewDelivery(
	logger *logrus.Entry,

	userHandler *user.Handler,

	cfg *config.Config,
) *Delivery {
	d := Delivery{
		logger: logger.WithFields(logrus.Fields{
			"layer":    "delivery",
			"delivery": "http",
		}),

		e: echo.New(),

		addr: fmt.Sprintf("%s:%d", cfg.Delivery.HTTP.Host, cfg.Delivery.HTTP.Port),

		userHandler: userHandler,
	}

	d.e.Logger.SetOutput(io.Discard)

	d.initRouter()

	return &d
}
