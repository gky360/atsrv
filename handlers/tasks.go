package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
	"github.com/labstack/echo"
	yaml "gopkg.in/yaml.v2"
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

	contestID, err := paramContest(c)
	if err != nil {
		return err
	}
	isFull := (c.QueryParam("full") == "true")

	page, err := getPage(h, user.ID)
	tasksPage, err := pages.NewTasksPage(page, contestID)
	if err != nil {
		return err
	}

	tasks, err := tasksPage.GetTasks()
	if err != nil {
		return err
	}
	if isFull {
		if err := getTasksFull(h, c, tasks); err != nil {
			return err
		}
	}

	rsp := new(RspGetTasks)
	rsp.Tasks = tasks

	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) GetTask(c echo.Context) (err error) {
	fmt.Println("h.GetTask")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, taskName, err := paramContestTask(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskName)

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

func paramContestTask(c echo.Context) (contestID, taskName string, err error) {
	contestID, err = paramContest(c)
	if err != nil {
		return
	}

	taskName = c.Param("taskName")
	if taskName == "" {
		err = echo.NewHTTPError(http.StatusBadRequest, "task name should not be empty.")
	}
	return
}

func getTasksFull(h *Handler, c echo.Context, tasks []models.Task) error {
	return nil
}
