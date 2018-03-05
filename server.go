package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"net/http"
	"os"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/pages"
)

func main() {
	e := echo.New()

	// user := new(models.User)

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		e.Logger.Fatal("Could not start chrome driver")
		e.Logger.Fatal(err)
		os.Exit(1)
	}
	defer driver.Stop()
	agoutiPage, err := driver.NewPage()
	if err != nil {
		e.Logger.Fatal("Could not create a page of chrome driver")
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "atsrv is running!")
	})

	e.POST("/login", func(c echo.Context) error {
		fmt.Println("trying to login ...")
		reqUser := new(models.User)
		if err := c.Bind(reqUser); err != nil {
			return err
		}

		page, err := pages.NewTasksPage(agoutiPage)
		if err != nil {
			return err
		}
		fmt.Println(page.GetPage().Title())

		resUser := *reqUser
		resUser.Password = ""
		return c.JSON(http.StatusOK, resUser)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
