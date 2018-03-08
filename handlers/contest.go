package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gky360/atsrv/models"
)

func (h *Handler) GetContest(c echo.Context) (err error) {
	fmt.Println("h.GetContest")

	contestID := c.Param("id")
	if len(contestID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "id should not be empty.")
	}

	// TODO: get contest
	contestFilePath := filepath.Join(h.PkgPath, "testdata", "contest.0.yaml")
	buf, err := ioutil.ReadFile(contestFilePath)
	if err != nil {
		panic(err)
	}
	contest := new(models.Contest)
	if err = yaml.Unmarshal(buf, &contest); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, contest)
}
