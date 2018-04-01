package handlers

import (
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"net/http"

	"github.com/gky360/atsrv/constants"
)

type (
	AtsrvConfig struct {
		UserID string `envconfig:"user_id"`
		Host   string `default:""`
		Port   string `default:"4700"`
	}

	Handler struct {
		page   *agouti.Page
		config AtsrvConfig
		token  string
	}

	RspRoot struct {
		Version string `json:"version"`
	}
)

func NewHandler(page *agouti.Page, config AtsrvConfig, token string) *Handler {
	return &Handler{
		page:   page,
		config: config,
		token:  token,
	}
}

func (h *Handler) Root(c echo.Context) error {
	rsp := new(RspRoot)
	rsp.Version = constants.Version
	return c.JSON(http.StatusOK, rsp)
}
