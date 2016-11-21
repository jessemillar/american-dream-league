package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetPositionByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetPositionByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllPositions(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllPositions()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdatePosition(context echo.Context) error {
	position := accessors.Position{}
	err := context.Bind(&position)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdatePosition(position)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakePosition(context echo.Context) error {
	position := accessors.Position{}
	err := context.Bind(&position)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakePosition(position)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeletePositionByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeletePositionByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
	return nil
}
