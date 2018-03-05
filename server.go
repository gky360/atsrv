package main

import (
	"fmt"
	"net/http"

	"github.com/gky360/atsrv/models"
	"github.com/labstack/echo"
)

func main() {
	// user := new(models.User)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "atsrv is running!")
	})

	e.POST("/login", func(c echo.Context) error {
		fmt.Println("trying to login ...")
		req_user := new(models.User)
		if err := c.Bind(req_user); err != nil {
			return err
		}

		// TODO: login

		res_user := *req_user
		res_user.Password = ""
		return c.JSON(http.StatusOK, res_user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
