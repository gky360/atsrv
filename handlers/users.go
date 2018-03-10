package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
)

func (h *Handler) Login(c echo.Context) (err error) {
	fmt.Println("h.Login")

	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return err
	}

	// TODO: validate user

	if err := startPage(h, u.ID); err != nil {
		return err
	}
	page, err := getPage(h, u.ID)
	if err != nil {
		return err
	}
	pageObj, err := pages.NewTasksPage(page)
	if err != nil {
		return err
	}
	fmt.Println(pageObj.GetPage().Title())

	// TODO: login

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Logout(c echo.Context) (err error) {
	fmt.Println("h.Logout")

	// TODO: logout
	// TODO: stop page

	u := new(models.User)
	return c.JSON(http.StatusOK, u)
}
