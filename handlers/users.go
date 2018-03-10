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

	if isLoggedIn(h, user.ID) {
		// when the user has already logged in, do not return token
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user %s has aready logged in.", user.ID))
	}

	page, err := startPage(h, user.ID)
	if err != nil {
		return err
	}
	pageObj, err := pages.NewTasksPage(page)
	if err != nil {
		stopPage(h, user.ID)
		return err
	}
	fmt.Println(pageObj.GetPage().Title())

	// TODO: access page
	// send user id and password

	// Generate encoded token and send it as response
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	user.Token, err = token.SignedString([]byte(h.jwtSecret))
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

func userIDFromToken(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func isLoggedIn(h *Handler, userID string) bool {
	_, err := getPage(h, userID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// TODO: access page
	// check if the user logged in to AtCoder

	return true
}

func currentUser(h *Handler, c echo.Context) (*models.User, error) {
	user := new(models.User)
	user.ID = userIDFromToken(c)
	loggedIn := isLoggedIn(h, user.ID)
	if !loggedIn {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("user %s is not logged in.", user.ID))
	}

	// TODO: access page
	// get user name
	user.Name = "myname"

	return user, nil
}
