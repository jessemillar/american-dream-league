package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetEmailByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetEmailByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllEmails(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllEmails()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateEmail(context echo.Context) error {
	email := accessors.Email{}
	err := context.Bind(&email)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdateEmail(email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeEmail(context echo.Context) error {
	email := accessors.Email{}
	err := context.Bind(&email)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeEmail(email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteEmailByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteEmailByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
}
