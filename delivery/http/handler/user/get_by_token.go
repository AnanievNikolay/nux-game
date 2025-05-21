package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/AnanievNikolay/nux-game/common/utils"
	handlerUtils "github.com/AnanievNikolay/nux-game/delivery/http/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/labstack/echo/v4"
)

// GetUserByToken godoc
// @Summary get user by token
// @Description get user by token
// @Tags user
// @Produce json
// @Param token path string true "token"
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /user/token/{token} [GET]
func (h *Handler) GetUserByToken(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	requestLogger := h.logger

	mf := utils.LogTimeSpent(requestLogger, "GetUserByToken")
	defer mf()

	token := c.Param("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "empty \"token\" param")
	}

	user, err := h.service.GetUserByToken(ctx, requestLogger, token)
	if err != nil {
		requestLogger.Errorf("service.CreateUser: %s", err)

		switch {
		case errors.Is(err, domain.ErrorUserNotFound):
			return c.NoContent(http.StatusNotFound)
		case errors.Is(err, domain.ErrTokenInvalidOrExpired):
			return c.String(http.StatusBadRequest, err.Error())
		default:
			return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
		}
	}

	return c.JSON(http.StatusOK, user)
}
