package token

import (
	"context"
	"errors"
	"net/http"

	"github.com/AnanievNikolay/nux-game/common/utils"
	handlerUtils "github.com/AnanievNikolay/nux-game/delivery/http/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/labstack/echo/v4"
)

// DeactivateToken godoc
// @Summary deactivate token
// @Description deactivate token
// @Tags token
// @Produce json
// @Param token path string true "token"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /game/{token}/deactivate [POST]
func (h *Handler) DeactivateToken(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	requestLogger := h.logger

	mf := utils.LogTimeSpent(requestLogger, "DeactivateToken")
	defer mf()

	token := c.Param("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "empty \"token\" param")
	}

	err := h.service.DeactivateToken(ctx, requestLogger, token)
	if err != nil {
		requestLogger.Errorf("service.DeactivateToken: %s", err)

		switch {
		case errors.Is(err, domain.ErrTokenInvalidOrExpired):
			return c.String(http.StatusBadRequest, err.Error())
		default:
			return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
		}
	}

	return c.NoContent(http.StatusOK)
}
