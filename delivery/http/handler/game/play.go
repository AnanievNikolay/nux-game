package game

import (
	"context"
	"errors"
	"net/http"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/delivery/http/handler/game/dto"
	handlerUtils "github.com/AnanievNikolay/nux-game/delivery/http/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/labstack/echo/v4"
)

// Play godoc
// @Summary play game
// @Description play game
// @Tags game
// @Produce json
// @Param token path string true "token"
// @Success 200 {object} dto.Game
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /game/{token}/play [POST]
func (h *Handler) Play(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	requestLogger := h.logger

	mf := utils.LogTimeSpent(requestLogger, "Play")
	defer mf()

	token := c.Param("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "empty \"token\" param")
	}

	gameResult, err := h.service.Play(ctx, requestLogger, token)
	if err != nil {
		requestLogger.Errorf("service.Play: %s", err)

		switch {
		case errors.Is(err, domain.ErrTokenInvalidOrExpired):
			return c.String(http.StatusBadRequest, err.Error())
		default:
			return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
		}
	}

	return c.JSON(http.StatusOK, dto.NewGame(gameResult))
}
