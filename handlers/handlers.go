package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"net/http"

	"github.com/gky360/atsrv/constants"
)

type (
	AtsrvConfig struct {
		UserID string
	}

	Handler struct {
		driver    *agouti.WebDriver
		userPages UserPages
		config    AtsrvConfig
	}

	UserPages map[string]*agouti.Page

	RspRoot struct {
		Version     string   `json:"version"`
		PageUserIDs []string `json:"page_user_ids"`
	}
)

func NewHandler(driver *agouti.WebDriver, config AtsrvConfig) *Handler {
	return &Handler{
		driver:    driver,
		userPages: UserPages{},
		config:    config,
	}
}

func (h *Handler) Root(c echo.Context) error {
	rsp := new(RspRoot)
	rsp.Version = constants.Version
	rsp.PageUserIDs = make([]string, len(h.userPages))
	i := 0
	for k := range h.userPages {
		rsp.PageUserIDs[i] = k
		i++
	}
	return c.JSON(http.StatusOK, rsp)
}

func startPage(h *Handler, userID string) (*agouti.Page, error) {
	if len(userID) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "could not start page for empty user id.")
	}
	if h.userPages[userID] == nil {
		page, err := h.driver.NewPage()
		if err != nil {
			return nil, err
		}
		h.userPages[userID] = page
	}
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
	if h.userPages[userID] != nil {
		if err := h.userPages[userID].Destroy(); err != nil {
			return err
		}
		delete(h.userPages, userID)
	}
	return nil
}
