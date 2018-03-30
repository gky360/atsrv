package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
)

func (h *Handler) Me(c echo.Context) (err error) {
	fmt.Println("h.Me")
	user, err := currentUserWithContestID(h, c, pages.PracticeContestID)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	return c.JSON(http.StatusOK, user)
}

func isLoggedIn(h *Handler, userID string, contestID string) bool {
	contestPage, err := getContestPage(h, userID, contestID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Check if the user logged in to AtCoder
	ret, err := contestPage.Navbar().IsLoggedIn()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return ret
}

func currentUserWithContestID(h *Handler, c echo.Context, contestID string) (*models.User, error) {
	user := new(models.User)
	user.ID = h.config.UserID
	loggedIn := isLoggedIn(h, user.ID, contestID)
	if !loggedIn {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("user %s is not logged in.", user.ID))
	}

	// Get user name
	contestPage, err := getContestPage(h, user.ID, contestID)
	if err != nil {
		return nil, err
	}
	userName, err := contestPage.Navbar().GetUserName()
	if err != nil {
		return nil, err
	}
	user.Name = userName

	return user, nil
}

func currentUser(h *Handler, c echo.Context) (*models.User, error) {
	contestID, err := paramContest(c)
	if err != nil {
		return nil, err
	}
	return currentUserWithContestID(h, c, contestID)
}
