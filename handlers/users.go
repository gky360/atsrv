package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
)

func (h *Handler) Login(c echo.Context) (err error) {
	fmt.Println("h.Login")

	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return err
	}

	if user.ID == "" || user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id or password.")
	}

	if isLoggedIn(h, user.ID, pages.PracticeContestID) {
		// when the user has already logged in, do not return token
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user %s has aready logged in.", user.ID))
	}

	page, err := startPage(h, user.ID)
	if err != nil {
		return err
	}
	loginPage, err := pages.NewLoginPage(page)
	if err != nil {
		stopPage(h, user.ID)
		return err
	}

	// Send user id and password
	if err := loginPage.Login(user.ID, user.Password); err != nil {
		return err
	}
	if !isLoggedIn(h, user.ID, pages.PracticeContestID) {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("failed to login to AtCoder as %s.", user.ID))
	}

	// Generate encoded token and send it as response
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	user.Token, err = token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		stopPage(h, user.ID)
		return err
	}

	// Get user name
	contestPage, err := pages.NewContestPage(page, pages.PracticeContestID)
	if err != nil {
		stopPage(h, user.ID)
		return err
	}
	user.Name, err = contestPage.Navbar().GetUserName()
	if err != nil {
		stopPage(h, user.ID)
		return err
	}

	user.Password = "" // Don't send password
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) Logout(c echo.Context) (err error) {
	fmt.Println("h.Logout")

	userID := userIDFromToken(c)

	if err = stopPage(h, userID); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.User{ID: userID})
}

func (h *Handler) Me(c echo.Context) (err error) {
	fmt.Println("h.Me")
	user, err := currentUserWithContestID(h, c, pages.PracticeContestID)
	if err != nil {
		return err
	}
	c.Logger().Info(user.ID)

	return c.JSON(http.StatusOK, user)
}

func userIDFromToken(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
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
	user.ID = userIDFromToken(c)
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
