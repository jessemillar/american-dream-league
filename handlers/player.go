package handlers

import (
	"net/http"

	"github.com/jessemillar/american-dream-league/helpers"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GenerateHitter(context echo.Context) error {
	response, err := helpers.GenerateHitter()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
