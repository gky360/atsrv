package handlers

import (
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"net/http"
)

type (
	Handler struct {
		PkgPath string
		Page    *agouti.Page
	}
)

func (h *Handler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "atsrv is running!")
}
