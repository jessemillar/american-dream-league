package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetPlayerByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetPlayerByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllPlayers(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllPlayers()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdatePlayer(context echo.Context) error {
	player := accessors.Player{}
	err := context.Bind(&player)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdatePlayer(player)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakePlayer(context echo.Context) error {
	player := accessors.Player{}
	err := context.Bind(&player)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakePlayer(player)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeletePlayerByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeletePlayerByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
	return nil
}
