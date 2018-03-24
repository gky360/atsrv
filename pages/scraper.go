package pages

import (
	"fmt"
	"path/filepath"

	"github.com/gky360/atsrv/constants"
	"github.com/sclevine/agouti"
)

const (
	PracticeContestID = "practice"
)

type Scraper interface {
	Page() *agouti.Page
	TargetPath() string
}

func TargetURL(s Scraper) string {
	return "https://" + filepath.Join(constants.AtCoderHost, s.TargetPath())
}

func At(s Scraper) (bool, error) {
	currentURL, err := s.Page().URL()
	if err != nil {
		return false, err
	}
	targetURL := TargetURL(s)
	return (currentURL == targetURL), nil
}

func To(s Scraper) error {
	isAt, err := At(s)
	if err != nil {
		return err
	}
	if !isAt {
		targetURL := TargetURL(s)
		fmt.Println(targetURL)
		if err := s.Page().Navigate(targetURL); err != nil {
			return err
		}
		isAt, err := At(s)
		if err != nil {
			return err
		}
		if !isAt {
			return fmt.Errorf("Couldn't navigate to %s", targetURL)
		}
	}
	return nil
}
