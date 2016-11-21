package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetTeamByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetTeamByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllTeams(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllTeams()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateTeam(context echo.Context) error {
	team := accessors.Team{}
	err := context.Bind(&team)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdateTeam(team)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeTeam(context echo.Context) error {
	team := accessors.Team{}
	err := context.Bind(&team)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeTeam(team)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteTeamByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteTeamByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
	return nil
}
