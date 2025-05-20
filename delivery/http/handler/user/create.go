package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/delivery/http/handler/user/dto"
	handlerUtils "github.com/AnanievNikolay/nux-game/delivery/http/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary create user
// @Description create user by username and phone
// @Tags user
// @Param request body dto.UserCreateRequest true "UserCreateRequest"
// @Produce json
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /user/register [POST]
func (h *Handler) CreateUser(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	requestLogger := h.logger

	mf := utils.LogTimeSpent(requestLogger, "CreateUser")
	defer mf()

	var request dto.UserCreateRequest
	if err := c.Bind(&request); err != nil {
		err = fmt.Errorf("error bind request: %w", err)
		requestLogger.Errorln(err)

		return c.String(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var sb strings.Builder
			for _, fieldErr := range validationErrors {
				sb.WriteString(
					fmt.Sprintf(
						"Field '%s' failed on the '%s' tag; ",
						fieldErr.Field(),
						fieldErr.Tag(),
					),
				)
			}

			return c.String(http.StatusBadRequest, sb.String())
		} else {
			requestLogger.Errorf("validate.Struct: %s", err)

			return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
		}
	}

	user, err := h.service.CreateUser(ctx, requestLogger, request.Username, request.Phone)
	if err != nil {
		requestLogger.Errorf("service.CreateUser: %s", err)

		if errors.Is(err, domain.ErrorUsernameWithThisPhoneNotUnique) {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.String(http.StatusInternalServerError, handlerUtils.InternalServerErrorMessage)
	}

	return c.JSON(http.StatusOK, user)
}
