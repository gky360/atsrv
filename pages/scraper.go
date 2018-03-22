package pages

import (
	"path/filepath"

	"github.com/gky360/atsrv/constants"
	"github.com/sclevine/agouti"
)

const (
	PracticeContestID = "practice"
)

type Scraper interface {
	Page() *agouti.Page
	TargetHost() string
	TargetPath() string
}

func ContestHost(contestID string) string {
	return "https://" + contestID + ".contest." + constants.AtCoderHost
}

func TargetURL(s Scraper) string {
	return filepath.Join(s.TargetHost(), s.TargetPath())
}

func At(s Scraper) error {
	page := s.Page()
	currentURL, err := page.URL()
	if err != nil {
		return err
	}
	targetURL := TargetURL(s)
	if currentURL != targetURL {
		if err := page.Navigate(targetURL); err != nil {
			return err
		}
	}
	return nil
}
