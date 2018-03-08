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

type (
	RspGetTasks struct {
		Tasks []models.Task `json:"tasks" yaml:"tasks"`
	}
)

func (h *Handler) GetTasks(c echo.Context) (err error) {
	fmt.Println("h.GetTasks")

	contestID, err := paramContestID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)

	// TODO: access page
	testFilePath := filepath.Join(h.PkgPath, "testdata", "tasks.yaml")
	buf, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	rsp := new(RspGetTasks)
	if err = yaml.Unmarshal(buf, &rsp); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, rsp)
}
