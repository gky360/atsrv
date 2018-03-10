package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"net/http"
)

type (
	Handler struct {
		pkgPath   string
		driver    *agouti.WebDriver
		jwtSecret []byte
		userPages UserPages
	}

	UserPages map[string]*agouti.Page
)

func NewHandler(pkgPath string, driver *agouti.WebDriver, jwtSecret []byte) *Handler {
	return &Handler{
		pkgPath:   pkgPath,
		driver:    driver,
		jwtSecret: jwtSecret,
		userPages: UserPages{},
	}
}

func (h *Handler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "atsrv is running!")
}

func startPage(h *Handler, userID string) (*agouti.Page, error) {
	if len(userID) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "could not start page for empty user id.")
	}
	if h.userPages[userID] != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("page for %s has already started.", userID))
	}
	page, err := h.driver.NewPage()
	if err != nil {
		return nil, err
	}
	h.userPages[userID] = page
	return h.userPages[userID], nil
}

func getPage(h *Handler, userID string) (*agouti.Page, error) {
	if len(userID) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "could not get page for empty user id.")
	}
	if h.userPages[userID] == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not get page for user %s", userID))
	}
	return h.userPages[userID], nil
}

func stopPage(h *Handler, userID string) error {
	if len(userID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "could not stop page for empty user id.")
	}
	if h.userPages[userID] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("page for %s is not running.", userID))
	}
	if err := h.userPages[userID].Destroy(); err != nil {
		return err
	}
	delete(h.userPages, userID)
	return nil
}
