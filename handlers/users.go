package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

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

func Login(h *Handler, password string) error {
	loginPage, err := pages.NewLoginPage(h.page)
	if err != nil {
		return err
	}

	// Send user id and password
	if err := loginPage.Login(h.config.UserID, password); err != nil {
		return err
	}
	if !isLoggedIn(h, pages.PracticeContestID) {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("failed to login to AtCoder as %s.", h.config.UserID))
	}

	return nil
}

func isLoggedIn(h *Handler, contestID string) bool {
	contestPage, err := getContestPage(h, contestID)
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
	loggedIn := isLoggedIn(h, contestID)
	if !loggedIn {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("user %s is not logged in.", user.ID))
	}

	// Get user name
	contestPage, err := getContestPage(h, contestID)
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
