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
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, err := paramContest(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)

	// TODO: access page
	testFilePath := filepath.Join(h.pkgPath, "testdata", "contest.yaml")
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

func paramContest(c echo.Context) (string, error) {
	contestID := c.Param("contestID")
	if contestID == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "contest id should not be empty.")
	}
	return contestID, nil
}
