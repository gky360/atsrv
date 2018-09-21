package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sclevine/agouti"

	"github.com/gky360/atsrv/constants"
)

type (
	AtsrvConfig struct {
		UserID    string `split_words:"true"`
		AuthToken string `split_words:"true"`
		Host      string `default:"localhost"`
		Port      string `default:"4700"`
		Headless  bool
		Debug     bool
	}

	Handler struct {
		page   *agouti.Page
		config AtsrvConfig
	}

	RspRoot struct {
		Version string `json:"version"`
	}
)

func NewHandler(page *agouti.Page, config AtsrvConfig) *Handler {
	return &Handler{
		page:   page,
		config: config,
	}
}

func (h *Handler) Root(c echo.Context) error {
	rsp := new(RspRoot)
	rsp.Version = constants.Version
	return c.JSON(http.StatusOK, rsp)
}
