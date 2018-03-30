package main

import (
	// "crypto/rand"
	"fmt"
	"os"

	"github.com/gky360/atsrv/handlers"
	"github.com/howeyc/gopass"
	"github.com/kelseyhightower/envconfig"
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

func run() int {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	var config handlers.AtsrvConfig
	if err := envconfig.Process("atsrv", &config); err != nil {
		e.Logger.Error(err)
		return 1
	}

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		e.Logger.Error("Could not start chromedriver")
		e.Logger.Error(err)
		return 1
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		e.Logger.Error("Could not open page in chromedriver")
		e.Logger.Error(err)
		return 1
	}
	defer page.Destroy()
	defer fmt.Println("Stopping server...")

	// get user id and password
	if config.UserID == "" {
		fmt.Print("AtCoder user id: ")
		fmt.Scan(&config.UserID)
	}
	fmt.Print("AtCoder password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		e.Logger.Error("Could not get password from input")
		e.Logger.Error(err)
		return 1
	}

	h := handlers.NewHandler(page, config)

	if err := handlers.Login(h, string(password)); err != nil {
		e.Logger.Error("Failed to login to AtCoder. Please make suer your user id and password are correct.")
		e.Logger.Error(err)
		return 1
	}

	// Middlewares
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", h.Root)

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

	return 0
}

func main() {
	os.Exit(run())
}
