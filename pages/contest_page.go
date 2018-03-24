package pages

import (
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
