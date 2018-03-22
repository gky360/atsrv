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

func (p *LoginPage) UserIDForm() *agouti.Selection {
	return p.page.First("input[name=name]")
}

func (p *LoginPage) PasswordForm() *agouti.Selection {
	return p.page.First("input[name=password]")
}

// Values

// Operations

func (p *LoginPage) Login(userID, password string) error {
	if err := p.UserIDForm().Fill(userID); err != nil {
		return err
	}
	if err := p.PasswordForm().Fill(password); err != nil {
		return err
	}
	if err := p.UserIDForm().Submit(); err != nil {
		return err
	}
	return nil
}
