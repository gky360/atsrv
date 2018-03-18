package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sclevine/agouti"
	"path/filepath"
	"runtime"

	"github.com/gky360/atsrv/handlers"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

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

	jwtSecret := []byte("hogehoge")
	h := handlers.NewHandler(exPath, driver, jwtSecret)

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/" || c.Path() == "/login" {
				// Skip authentication for root and login requests
				return true
			}
			return false
		},
	}))

	// Routes
	e.GET("/", h.Root)

	e.POST("/login", h.Login)
	e.POST("/logout", h.Logout)
	e.GET("/me", h.Me)

	e.GET("/contests/:contestID", h.GetContest)

	e.GET("/contests/:contestID/tasks", h.GetTasks)
	e.GET("/contests/:contestID/tasks/:taskName", h.GetTask)

	e.GET("/contests/:contestID/submissions", h.GetSubmissions)
	e.GET("/contests/:contestID/submissions/:submissionID", h.GetSubmission)
	e.POST("/contests/:contestID/submissions", h.PostSubmission)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
