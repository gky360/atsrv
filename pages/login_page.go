package pages

import (
	"github.com/sclevine/agouti"
)

type LoginPage struct {
	page      *agouti.Page
	contestID string
}

func (p *LoginPage) Page() *agouti.Page {
	return p.page
}

func (p *LoginPage) TargetHost() string {
	return ContestHost(p.contestID)
}

func (p *LoginPage) TargetPath() string {
	return "/login"
}

func NewLoginPage(page *agouti.Page) (*LoginPage, error) {
	p := &LoginPage{
		page:      page,
		contestID: PracticeContestID,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *LoginPage) userIDForm() *agouti.Selection {
	return p.page.Find("input[name=name]")
}

func (p *LoginPage) passwordForm() *agouti.Selection {
	return p.page.Find("input[name=password]")
}

// Values

// Funcs

func (p *LoginPage) Login(userID, password string) error {
	if err := p.userIDForm().Fill(userID); err != nil {
		return err
	}
	if err := p.passwordForm().Fill(password); err != nil {
		return err
	}
	if err := p.userIDForm().Submit(); err != nil {
		return err
	}
	return nil
}
