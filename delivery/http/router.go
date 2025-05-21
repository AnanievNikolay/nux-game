package http

import (
	nethttp "net/http"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/AnanievNikolay/nux-game/docs"
)

// imported dto for openAPI Docs
func (d *Delivery) initRouter() {
	// CORS
	d.e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{
					nethttp.MethodGet, nethttp.MethodHead, nethttp.MethodPut,
					nethttp.MethodPatch, nethttp.MethodPost, nethttp.MethodDelete,
				},
			},
		))

	// is_alive
	d.e.GET("/is_alive", d.check)

	d.e.GET("/docs/*", echoSwagger.WrapHandler)

	d.e.POST("/user/register", d.userHandler.CreateUser)

	d.e.GET("/user/token/:token", d.userHandler.GetUserByToken)
}
