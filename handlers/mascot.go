package handlers

import (
	"net/http"
	"strconv"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetMascotByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetMascotByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetAllMascots(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllMascots()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateMascot(context echo.Context) error {
	mascot := accessors.Mascot{}
	err := context.Bind(&mascot)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.UpdateMascot(mascot)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeMascot(context echo.Context) error {
	mascot := accessors.Mascot{}
	err := context.Bind(&mascot)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeMascot(mascot)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) DeleteMascotByID(context echo.Context) error {
	ID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	err = handlerGroup.Accessors.DeleteMascotByID(ID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	jsonresp.New(context.Response(), http.StatusOK, "Successfully deleted")
	return nil
}
