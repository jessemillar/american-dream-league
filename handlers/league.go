package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetLeagueByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetLeagueByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllLeagues(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllLeagues()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateLeague(context echo.Context) error {
	response, err := handlerGroup.Accessors.UpdateLeague(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeLeague(context echo.Context) error {
	response, err := handlerGroup.Accessors.MakeLeague(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteLeagueByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteLeagueByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, "")
}
