package token

import (
	"context"
	"errors"
	"net/http"

	"github.com/AnanievNikolay/nux-game/common/utils"
	handlerUtils "github.com/AnanievNikolay/nux-game/delivery/http/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/AnanievNikolay/nux-game/repository/token/sqlite/dto"
	"github.com/labstack/echo/v4"
)

// UpdateToken godoc
// @Summary update token. Deactivate old and issue new
// @Description update token
// @Tags token
// @Produce json
// @Param token path string true "token"
// @Success 200 {object} dto.Token
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /game/{token}/renew [POST]
func (h *Handler) UpdateToken(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	requestLogger := h.logger

	mf := utils.LogTimeSpent(requestLogger, "UpdateToken")
	defer mf()

	token := c.Param("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "empty \"token\" param")
	}

	newToken, err := h.service.UpdateToken(ctx, requestLogger, token)
	if err != nil {
		requestLogger.Errorf("service.UpdateToken: %s", err)

		switch {
		case errors.Is(err, domain.ErrTokenInvalidOrExpired):
			return c.String(http.StatusBadRequest, err.Error())
		default:
			return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
		}
	}

	return c.JSON(http.StatusOK, dto.NewToken(newToken))
}
