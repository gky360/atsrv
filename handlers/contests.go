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

	contestID, err := paramContestID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)

	// TODO: access page
	testFilePath := filepath.Join(h.PkgPath, "testdata", "contest.yaml")
	buf, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	contest := new(models.Contest)
	if err = yaml.Unmarshal(buf, &contest); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, contest)
}
