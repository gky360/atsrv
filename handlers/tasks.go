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
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, err := paramContestID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)

	// TODO: access page
	testFilePath := filepath.Join(h.pkgPath, "testdata", "tasks.yaml")
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

func (h *Handler) GetTask(c echo.Context) (err error) {
	fmt.Println("h.GetTask")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, taskID, err := paramContestTaskID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskID)

	// TODO: access page
	testFilePath := filepath.Join(h.pkgPath, "testdata", "tasks.yaml")
	buf, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	rsp := new(RspGetTasks)
	if err = yaml.Unmarshal(buf, &rsp); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, rsp.Tasks[0])
}

func paramContestTaskID(c echo.Context) (contestID, taskID string, err error) {
	contestID, err = paramContestID(c)
	if err != nil {
		return
	}

	taskID = c.Param("taskID")
	if len(taskID) == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, "task id should not be empty.")
	}
	return
}