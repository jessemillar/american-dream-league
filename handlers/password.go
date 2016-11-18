package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetPasswordByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetPasswordByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllPasswords(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllPasswords()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdatePassword(context echo.Context) error {
	response, err := handlerGroup.Accessors.UpdatePassword(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakePassword(context echo.Context) error {
	response, err := handlerGroup.Accessors.MakePassword(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeletePasswordByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeletePasswordByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return jsonresp.New(context, http.StatusOK, "Successfully deleted")
}
