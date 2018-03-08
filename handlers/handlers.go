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

func paramContestID(c echo.Context) (string, error) {
	contestID := c.Param("contestID")
	if len(contestID) == 0 {
		return "", echo.NewHTTPError(http.StatusBadRequest, "id should not be empty.")
	}
	return contestID, nil
}
