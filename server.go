package main

import (
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"

	"github.com/gky360/atsrv/handlers"
)

func main() {
	e := echo.New()

	// user := new(models.User)

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		e.Logger.Error("Could not start chrome driver")
		e.Logger.Fatal(err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		e.Logger.Error("Could not create a page of chrome driver")
		e.Logger.Fatal(err)
	}

	h := &handlers.Handler{
		Page: page,
	}

	// Routes
	e.GET("/", h.Root)
	e.POST("/login", h.Login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
