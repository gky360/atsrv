package handlers

import (
	"github.com/sclevine/agouti"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	Page *agouti.Page
}

func (h *Handler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "atsrv is running!")
}
