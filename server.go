package main

import (
	"crypto/rand"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/gky360/atsrv/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sclevine/agouti"
)

const (
	version = "0.0.1"
)

var (
	banner = fmt.Sprintf(`
        __
       /\ \__
   __  \ \ ,_\   ____  _ __   __  __
 /'__'\ \ \ \/  /',__\/\''__\/\ \/\ \
/\ \L\.\_\ \ \_/\__, '\ \ \/ \ \ \_/ |
\ \__/.\_\\ \__\/\____/\ \_\  \ \___/
 \/__/\/_/ \/__/\/___/  \/_/   \/__/
%38s
`, "v"+version)
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
		return
	}
	defer driver.Stop()

	// generate secret for json web token
	jwtSecret := make([]byte, 16)
	_, err := rand.Read(jwtSecret)
	if err != nil {
		e.Logger.Error("Could not generate server secret")
		e.Logger.Fatal(err)
		return
	}
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
	e.POST("/contests/:contestID/join", h.Join)

	e.GET("/contests/:contestID/tasks", h.GetTasks)
	e.GET("/contests/:contestID/tasks/:taskName", h.GetTask)

	e.GET("/contests/:contestID/submissions", h.GetSubmissions)
	e.GET("/contests/:contestID/submissions/:submissionID", h.GetSubmission)
	e.POST("/contests/:contestID/submissions", h.PostSubmission)

	// Start server
	e.HideBanner = true
	fmt.Println(banner)
	e.Logger.Fatal(e.Start(":1323"))
}
