package views

import (
	"net/http"

	"github.com/labstack/echo"
)

func Frontend(context echo.Context) error {
	return context.Render(http.StatusOK, "frontend", "")
}
