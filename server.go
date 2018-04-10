package main

import (
	"crypto/rand"
	"encoding/base64"
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

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

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

	token, err := generateRandomString(8)
	if err != nil {
		e.Logger.Error("Could not generate server secret")
		e.Logger.Error(err)
		return 1
	}
	h := handlers.NewHandler(page, config, token)

	if err := handlers.Login(h, string(password)); err != nil {
		e.Logger.Error("Failed to login to AtCoder. Please make suer your user id and password are correct.")
		e.Logger.Error(err)
		return 1
	}

	// Middlewares
	e.Use(middleware.Logger())

	basicAuthConfig := middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/" {
				// Skip authentication for root endpoint
				return true
			}
			return false
		},
		Validator: func(authUserID, authToken string, c echo.Context) (bool, error) {
			if authUserID == config.UserID && authToken == token {
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

	fmt.Println("AuthToken:")
	fmt.Println(token)
	fmt.Println()

	e.Logger.Fatal(e.Start(config.Host + ":" + config.Port))

	return 0
}

func main() {
	os.Exit(run())
}
