package handler

import (
	"github.com/labstack/echo/v4"
	"github/patrickchap/booktalk/view"
)

type HomeHandler struct{}

func (h *HomeHandler) Index(c echo.Context) error {
	return render(c, view.Home())

}
