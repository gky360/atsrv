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
	RspGetSubmissions struct {
		Submissions []models.Submission `json:"submissions" yaml:"submissions"`
	}
)

func (h *Handler) GetSubmissions(c echo.Context) (err error) {
	fmt.Println("h.GetSubmissions")
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
	testFilePath := filepath.Join(h.pkgPath, "testdata", "submissions.yaml")
	buf, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	rsp := new(RspGetSubmissions)
	if err = yaml.Unmarshal(buf, &rsp); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) GetSubmission(c echo.Context) (err error) {
	fmt.Println("h.GetSubmission")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, taskName, submissionID, err := paramContestTaskSubmission(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskName)
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

	contestID, taskName, err := paramContestTask(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskName)
	sbm := new(models.Submission)
	if err = c.Bind(sbm); err != nil {
		return err
	}
	if len(sbm.Source) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "source should not be empty.")
	}

	// TODO: access page

	return c.JSON(http.StatusOK, sbm)
}

func paramContestTaskSubmission(c echo.Context) (
	contestID, taskName, SubmissionID string,
	err error,
) {
	contestID, taskName, err = paramContestTask(c)
	if err != nil {
		return
	}

	submissionID := c.Param("submissionID")
	if len(submissionID) == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, "submission id should not be empty.")
	}
	return
}
