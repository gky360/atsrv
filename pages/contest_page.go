package pages

import (
	"fmt"
	"strings"

	"github.com/gky360/atsrv/models"
	"github.com/gky360/atsrv/modules"
	"github.com/sclevine/agouti"
)

type ContestPage struct {
	page      *agouti.Page
	contestID string
	navbar    *modules.NavbarModule
}

func (p *ContestPage) Page() *agouti.Page {
	return p.page
}

func (p *ContestPage) TargetHost() string {
	return ContestHost(p.contestID)
}

func (p *ContestPage) TargetPath() string {
	return "/#"
}

func NewContestPage(page *agouti.Page, contestID string) (*ContestPage, error) {
	p := &ContestPage{
		page:      page,
		contestID: contestID,
		navbar:    modules.NewNavbarModule(page),
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *ContestPage) Navbar() *modules.NavbarModule {
	return p.navbar
}

// Values

func (p *ContestPage) contestName() (string, error) {
	const selector = ".insert-participant-box h1"
	rawStr, err := p.page.Find(selector).Text()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(rawStr), nil
}

// Funcs

func (p *ContestPage) GetContest() (*models.Contest, error) {
	contestName, err := p.contestName()
	if err != nil {
		return nil, err
	}
	return &models.Contest{
		ID:   p.contestID,
		Name: contestName,
	}, nil
}
