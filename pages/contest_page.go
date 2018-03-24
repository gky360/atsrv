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

func (p *ContestPage) joinBtn() *agouti.Selection {
	const selector = ".insert-participant-box button"
	return p.page.Find(selector)
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

func (p *ContestPage) IsJoined() (bool, error) {
	cnt, _ := p.joinBtn().Count()
	if cnt >= 2 {
		return false, fmt.Errorf("found multiple join buttons")
	} else if cnt == 1 {
		// not joined
		return false, nil
	}
	// joined
	return true, nil
}

func (p *ContestPage) Join() error {
	isJoined, err := p.IsJoined()
	if err != nil {
		return err
	}
	if isJoined {
		return nil
	}

	if err := p.joinBtn().Click(); err != nil {
		return err
	}
	isJoined, err = p.IsJoined()
	if err != nil {
		return err
	}
	if !isJoined {
		return fmt.Errorf("failed to join contest %s", p.contestID)
	}
	return nil
}
