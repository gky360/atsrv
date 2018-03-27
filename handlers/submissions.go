package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
	"github.com/labstack/echo"
	yaml "gopkg.in/yaml.v2"
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

	page, err := getPage(h, user.ID)
	if err != nil {
		return err
	}
	taskID := ""
	if taskName != "" {
		tasksPage, err := pages.NewTasksPage(page, contestID)
		if err != nil {
			return err
		}
		taskID, err = tasksPage.GetTaskID(taskName)
		if err != nil {
			return err
		}
	}
	sbmsPage, err := pages.NewSubmissionsPage(page, contestID, taskID, models.LangNone)
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

	contestID, submissionID, err := paramContestSubmission(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(submissionID)

	// TODO: access page
	testFilePath := filepath.Join(h.pkgPath, "testdata", "submissions.yaml")
	buf, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	rsp := new(RspGetSubmissions)
	if err = yaml.Unmarshal(buf, &rsp); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, rsp.Submissions[0])
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
	fmt.Println(contestID)
	fmt.Println(taskName)
	sbm := new(models.Submission)
	if err = c.Bind(sbm); err != nil {
		return err
	}
	if sbm.Source == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "source should not be empty.")
	}

	// TODO: access page

	return c.JSON(http.StatusOK, sbm)
}

func paramContestTaskQ(c echo.Context) (contestID, taskName string, err error) {
	contestID, err = paramContest(c)
	if err != nil {
		return
	}

	taskName = c.QueryParam("task_name")
	return
}

func paramContestSubmission(c echo.Context) (contestID string, submissionID int, err error) {
	contestID, err = paramContest(c)
	if err != nil {
		return
	}

	submissionID, err = strconv.Atoi(c.Param("submissionID"))
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "submission id is invalid.")
	}
	return
}
