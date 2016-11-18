package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetNameByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetNameByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllNames(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllNames()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateName(context echo.Context) error {
	name := accessors.Name{}
	err := context.Bind(&name)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdateName(name)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeName(context echo.Context) error {
	name := accessors.Name{}
	err := context.Bind(&name)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeName(name)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteNameByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteNameByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
}
