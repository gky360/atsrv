package pages

import (
	"github.com/sclevine/agouti"
	"path"

	"github.com/gky360/atsrv/models"
)

type Page interface {
	GetTargetPath() string
	GetPage() *agouti.Page
}

func GetTargetURL(p Page) string {
	contest, _ := models.GetCurrentContest()
	return path.Join(contest.URL(), p.GetTargetPath())
}

func At(p Page) error {
	page := p.GetPage()
	currentURL, err := page.URL()
	if err != nil {
		return err
	}
	targetURL := GetTargetURL(p)
	if currentURL != targetURL {
		if err := page.Navigate(targetURL); err != nil {
			return err
		}
	}
	return nil
}
