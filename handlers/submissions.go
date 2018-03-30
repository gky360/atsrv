package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
	"github.com/labstack/echo"
)

type (
	RspGetSubmissions struct {
		Submissions []*models.Submission `json:"submissions" yaml:"submissions"`
	}
)

func (h *Handler) GetSubmissions(c echo.Context) (err error) {
	fmt.Println("h.GetSubmissions")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, taskName, err := paramContestTaskQ(c)
	if err != nil {
		return err
	}

	taskID := ""
	if taskName != "" {
		tasksPage, err := pages.NewTasksPage(h.page, contestID)
		if err != nil {
			return err
		}
		taskID, err = tasksPage.GetTaskID(taskName)
		if err != nil {
			return err
		}
	}
	sbmsPage, err := pages.NewSubmissionsPage(h.page, contestID, taskID, models.LangNone)
	if err != nil {
		return err
	}

	sbms, err := sbmsPage.GetSubmissions()
	if err != nil {
		return err
	}

	rsp := new(RspGetSubmissions)
	rsp.Submissions = sbms

	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) GetSubmission(c echo.Context) (err error) {
	fmt.Println("h.GetSubmission")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, sbmID, err := paramContestSubmission(c)
	if err != nil {
		return err
	}

	sbmsPage, err := pages.NewSubmissionsPage(h.page, contestID, "", models.LangNone)
	if err != nil {
		return err
	}

	sbm, err := sbmsPage.GetSubmission(sbmID)
	if err != nil {
		if err == pages.ErrSubmissionNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("could not find submission with id %d", sbmID))
		}
		return err
	}

	return c.JSON(http.StatusOK, sbm)
}

func (h *Handler) PostSubmission(c echo.Context) (err error) {
	fmt.Println("h.PostSubmission")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, taskName, err := paramContestTaskQ(c)
	if err != nil {
		return err
	}
	if taskName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "task name should not be empty.")
	}
	sbm := new(models.Submission)
	if err = c.Bind(sbm); err != nil {
		return err
	}
	if sbm.Source == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "source should not be empty.")
	}

	tasksPage, err := pages.NewTasksPage(h.page, contestID)
	if err != nil {
		return err
	}
	taskID, err := tasksPage.GetTaskID(taskName)
	if err != nil {
		return err
	}
	taskPage, err := pages.NewTaskPage(h.page, contestID, taskID)
	if err != nil {
		return err
	}
	if err = taskPage.Submit(models.LangCpp14GCC, sbm.Source); err != nil {
		return err
	}

	sbmsPage, err := pages.NewSubmissionsPage(h.page, contestID, taskID, models.LangNone)
	if err != nil {
		return err
	}
	sbms, err := sbmsPage.GetSubmissions()
	if err != nil {
		return err
	}

	if len(sbms) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to submit source code")
	}

	return c.JSON(http.StatusOK, sbms[0])
}

func paramContestTaskQ(c echo.Context) (contestID, taskName string, err error) {
	contestID, err = paramContest(c)
	if err != nil {
		return
	}

	taskName = c.QueryParam("task_name")
	return
}

func paramContestSubmission(c echo.Context) (contestID string, sbmID int, err error) {
	contestID, err = paramContest(c)
	if err != nil {
		return
	}

	sbmID, err = strconv.Atoi(c.Param("submissionID"))
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "submission id is invalid.")
	}
	return
}
