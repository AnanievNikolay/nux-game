package http

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (d *Delivery) check(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
