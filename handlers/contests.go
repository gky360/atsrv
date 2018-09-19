package handlers

import (
	"fmt"
	"net/http"

	"github.com/gky360/atsrv/pages"
	"github.com/labstack/echo"
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
	withTestcasesURL := (c.QueryParam("with_testcases_url") == "true")

	testcasesURL := ""
	if withTestcasesURL {
		testcasesPage, err := getTestcasesPage(h)
		if err != nil {
			return err
		}
		testcasesURL, err = testcasesPage.GetContestFolderURL(contestID)
		if err != nil {
			if err == pages.ErrTestcasesFolderNotFound {
				return echo.NewHTTPError(http.StatusUnprocessableEntity,
					fmt.Sprintf("could not find testcases folder for contest %s", contestID))
			}
			return err
		}
	}

	contestPage, err := getContestPage(h, contestID)
	if err != nil {
		return err
	}
	contest, err := contestPage.GetContest()
	if err != nil {
		return err
	}
	if withTestcasesURL {
		contest.TestcasesURL = testcasesURL
	}

	return c.JSON(http.StatusOK, contest)
}

func (h *Handler) Join(c echo.Context) (err error) {
	fmt.Println("h.Join")
	user, err := currentUser(h, c)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	contestID, err := paramContest(c)
	if err != nil {
		return err
	}

	contestPage, err := getContestPage(h, contestID)
	if err != nil {
		return err
	}
	isJoined, err := contestPage.IsJoined()
	if err != nil {
		return err
	}
	if isJoined {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("already joined contest %s", contestID))
	}

	if err = contestPage.Join(); err != nil {
		return err
	}
	contest, err := contestPage.GetContest()
	if err != nil {
		return err
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

func getContestPage(h *Handler, contestID string) (*pages.ContestPage, error) {
	return pages.NewContestPage(h.page, contestID)
}

func getTestcasesPage(h *Handler) (*pages.TestcasesPage, error) {
	return pages.NewTestcasesPage(h.page)
}
