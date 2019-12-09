package handlers

import (
	"fmt"
	"net/http"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
	"github.com/labstack/echo/v4"
	"github.com/sclevine/agouti"
)

type (
	RspGetTasks struct {
		Tasks []*models.Task `json:"tasks" yaml:"tasks"`
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

	tasksPage, err := pages.NewTasksPage(h.page, contestID)
	if err != nil {
		return err
	}

	tasks, err := tasksPage.GetTasks()
	if err != nil {
		return err
	}
	if isFull {
		if err := getTasksFull(h.page, contestID, tasks); err != nil {
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

	tasksPage, err := pages.NewTasksPage(h.page, contestID)
	if err != nil {
		return err
	}
	taskID, err := tasksPage.GetTaskID(taskName)
	if err != nil {
		if err == pages.ErrTaskNameNotFound {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("could not find task name %s", taskName))
		}
		return err
	}
	taskPage, err := pages.NewTaskPage(h.page, contestID, taskID)
	if err != nil {
		return err
	}

	task, err := taskPage.GetTask()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
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

func getTasksFull(page *agouti.Page, contestID string, tasks []*models.Task) error {
	for i := range tasks {
		taskPage, err := pages.NewTaskPage(page, contestID, tasks[i].ID)
		if err != nil {
			return err
		}
		tasks[i], err = taskPage.GetTask()
		if err != nil {
			return err
		}
	}
	return nil
}
