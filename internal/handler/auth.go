package handler

import (
	"net/http"

	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) signUp(ctx echo.Context) error {
	var inputUser model.User

	err := ctx.Bind(&inputUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(inputUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, echo.Map{
		"id" : id,
	})

	return nil
}

func (h *Handler) signIn(ctx echo.Context) error {
	return nil
}