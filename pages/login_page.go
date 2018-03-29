package pages

import (
	"github.com/sclevine/agouti"
)

type LoginPage struct {
	page *agouti.Page
}

func (p *LoginPage) Page() *agouti.Page {
	return p.page
}

func (p *LoginPage) TargetPath() string {
	return "/login"
}

func NewLoginPage(page *agouti.Page) (*LoginPage, error) {
	p := &LoginPage{
		page: page,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *LoginPage) userIDForm() *agouti.Selection {
	const selector = "input#username"
	return p.page.Find(selector)
}

func (p *LoginPage) passwordForm() *agouti.Selection {
	const selector = "input#password"
	return p.page.Find(selector)
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
