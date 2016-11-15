package handlers

import (
	"net/http"

	"github.com/jessemillar/american-dream-league/helpers"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetHitters(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetHitters()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetPitchers(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetPitchers()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GenerateHitter(context echo.Context) error {
	response, err := helpers.GenerateHitter()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GeneratePitcher(context echo.Context) error {
	response, err := helpers.GeneratePitcher()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
