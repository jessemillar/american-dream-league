package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetUserByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetUserByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllUsers(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllUsers()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateUser(context echo.Context) error {
	user := accessors.User{}
	err := context.Bind(&user)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdateUser(user)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeUser(context echo.Context) error {
	user := accessors.User{}
	err := context.Bind(&user)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeUser(user)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteUserByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteUserByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return jsonresp.New(context, http.StatusOK, "Successfully deleted")
}
