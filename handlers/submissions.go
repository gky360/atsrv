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

	contestID, taskID, err := paramContestTaskID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskID)

	// TODO: access page
	testFilePath := filepath.Join(h.PkgPath, "testdata", "submissions.yaml")
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

	contestID, taskID, submissionID, err := paramContestTaskSubmissionID(c)
	if err != nil {
		return err
	}
	fmt.Println(contestID)
	fmt.Println(taskID)
	fmt.Println(submissionID)

	// TODO: access page
	testFilePath := filepath.Join(h.PkgPath, "testdata", "submissions.yaml")
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

func paramContestTaskSubmissionID(c echo.Context) (
	contestID, taskID, SubmissionID string,
	err error,
) {
	contestID, taskID, err = paramContestTaskID(c)
	if err != nil {
		return
	}

	submissionID := c.Param("submissionID")
	if len(submissionID) == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, "submission id should not be empty.")
	}
	return
}
