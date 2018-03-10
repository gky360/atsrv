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
		userPages UserPages
	}

	UserPages map[string]*agouti.Page
)

func NewHandler(pkgPath string, driver *agouti.WebDriver) *Handler {
	return &Handler{
		pkgPath:   pkgPath,
		driver:    driver,
		userPages: UserPages{},
	}
}

func (h *Handler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "atsrv is running!")
}

func startPage(h *Handler, userID string) error {
	if len(userID) == 0 {
		return fmt.Errorf("could not start page for empty user id.")
	}
	if h.userPages[userID] == nil {
		page, err := h.driver.NewPage()
		if err != nil {
			return err
		}
		h.userPages[userID] = page
	}
	return nil
}

func getPage(h *Handler, userID string) (*agouti.Page, error) {
	if h.userPages[userID] == nil {
		return nil, fmt.Errorf("could not get page for user %s.", userID)
	}
	return h.userPages[userID], nil
}

func stopPage(h *Handler, userID string) error {
	if h.userPages[userID] != nil {
		if err := h.userPages[userID].Destroy(); err != nil {
			return err
		}
	}
	return nil
}
