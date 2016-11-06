package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetLeagueById(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetLeagueById(context.Param("id"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetLeagueByName(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetLeagueByName(context.Param("name"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeLeague(context echo.Context) error {
	response, err := handlerGroup.Accessors.MakeLeague(context.Param("name"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
