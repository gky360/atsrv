package main

import (
	"github.com/labstack/echo"
	"github.com/sclevine/agouti"
	"path/filepath"
	"runtime"

	"github.com/gky360/atsrv/handlers"
)

func main() {
	e := echo.New()

	_, ex, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	exPath := filepath.Dir(ex)

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		e.Logger.Error("Could not start chrome driver")
		e.Logger.Fatal(err)
	}
	defer driver.Stop()

	h := handlers.NewHandler(exPath, driver)

	// Routes
	e.GET("/", h.Root)

	e.POST("/login", h.Login)
	e.POST("/logout", h.Logout)

	e.GET("/contests/:contestID", h.GetContest)

	e.GET("/contests/:contestID/tasks", h.GetTasks)
	e.GET("/contests/:contestID/tasks/:taskID", h.GetTask)

	e.GET("/contests/:contestID/tasks/:taskID/submissions", h.GetSubmissions)
	e.GET("/contests/:contestID/tasks/:taskID/submissions/:submissionID", h.GetSubmission)
	e.POST("/contests/:contestID/tasks/:taskID/submissions", h.PostSubmission)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
