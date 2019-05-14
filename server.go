package main

import (
	"fmt"
	"os"

	"github.com/gky360/atsrv/constants"
	"github.com/gky360/atsrv/handlers"
	"github.com/howeyc/gopass"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sclevine/agouti"
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
`, constants.Version)
)

func run() int {
	e := echo.New()

	var config handlers.AtsrvConfig
	if err := envconfig.Process("atsrv", &config); err != nil {
		e.Logger.Error(err)
		return 1
	}
	if config.Debug {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}

	// get auth token, user id and password
	if config.AuthToken == "" {
		e.Logger.Error("Set auth token to environment variable ATSRV_AUTH_TOKEN.\n" +
			"ATSRV_AUTH_TOKEN is used to communicate with atcli.")
		return 1
	}
	fmt.Print("AtCoder user id : ")
	if config.UserID == "" {
		fmt.Scan(&config.UserID)
	} else {
		fmt.Println(config.UserID)
	}
	fmt.Print("AtCoder password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		e.Logger.Error("Could not get password from input")
		e.Logger.Error(err)
		return 1
	}

	chromeOptions := []agouti.Option{}
	if config.Headless {
		chromeOptions = append(chromeOptions, agouti.ChromeOptions(
			"args",
			[]string{
				"--headless",
				"--disable-gpu",
			}),
		)
	}
	if config.Debug {
		chromeOptions = append(chromeOptions, agouti.Debug)
	}
	driver := agouti.ChromeDriver(chromeOptions...)
	if err := driver.Start(); err != nil {
		e.Logger.Error("Could not start chromedriver")
		e.Logger.Error(err)
		return 1
	}
	defer func() { _ = driver.Stop() }()
	page, err := driver.NewPage()
	if err != nil {
		e.Logger.Error("Could not open page in chromedriver")
		e.Logger.Error(err)
		return 1
	}
	defer func() { _ = page.Destroy() }()
	defer fmt.Println("Stopping server...")

	h := handlers.NewHandler(page, config)

	if err := handlers.Login(h, string(password)); err != nil {
		e.Logger.Error("Failed to login to AtCoder. Please make suer your user id and password are correct.")
		e.Logger.Error(err)
		return 1
	}

	// Middlewares
	e.Use(middleware.Logger())

	basicAuthConfig := middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			// Skip authentication for root endpoint
			return c.Path() == "/"
		},
		Validator: func(authUserID, authToken string, c echo.Context) (bool, error) {
			if authUserID == config.UserID && authToken == config.AuthToken {
				return true, nil
			}
			return false, nil
		},
	}
	e.Use(middleware.BasicAuthWithConfig(basicAuthConfig))

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
	fmt.Println()

	e.Logger.Fatal(e.Start(config.Host + ":" + config.Port))

	return 0
}

func main() {
	os.Exit(run())
}
