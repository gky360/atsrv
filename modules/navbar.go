package modules

import (
	"regexp"

	"github.com/sclevine/agouti"
)

type NavbarModule struct {
	page *agouti.Page
}

func NewNavbarModule(page *agouti.Page) *NavbarModule {
	return &NavbarModule{page: page}
}

// Elements

func (m *NavbarModule) self() *agouti.Selection {
	const navbarSelector = "div.navbar"
	return m.page.Find(navbarSelector)
}

func (m *NavbarModule) loginedUL() *agouti.Selection {
	const loginedUlSelector = ".nav-collapse #nav-right-logined"
	return m.self().Find(loginedUlSelector)
}

// Values

func (m *NavbarModule) userName() (string, error) {
	const userNameSelector = "#nav-right-username"
	userNameRaw, err := m.loginedUL().Find(userNameSelector).Text()
	if err != nil {
		return "", err
	}
	rep := regexp.MustCompile(`^(.*)\((?:contestant|guest)\)$`)
	return rep.ReplaceAllString(userNameRaw, "$1"), nil
}

// Funcs

func (m *NavbarModule) IsLoggedIn() (bool, error) {
	return m.loginedUL().Visible()
}

func (m *NavbarModule) GetUserName() (string, error) {
	isLoggedIn, err := m.IsLoggedIn()
	if err != nil {
		return "", err
	}
	if !isLoggedIn {
		return "", nil
	}
	return m.userName()
}
