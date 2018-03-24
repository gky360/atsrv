package modules

import (
	"regexp"
	"strings"

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
	const navbarSelector = "nav.navbar"
	return m.page.Find(navbarSelector)
}

func (m *NavbarModule) userDropdown() *agouti.Selection {
	const xpath = `//*[@id='navbar-collapse']/*[contains(@class, 'navbar-right')]/li[contains(@class, 'dropdown')][2]`
	return m.self().FindByXPath(xpath)
}

func (m *NavbarModule) userLink() *agouti.Selection {
	const selector = "a.dropdown-toggle"
	return m.userDropdown().Find(selector)
}

// Values

func (m *NavbarModule) userName() (string, error) {
	userNameRaw, err := m.userLink().Text()
	if err != nil {
		return "", err
	}
	rep := regexp.MustCompile(`^(.*) \((?:Contestant|Guest)\)$`)
	return rep.ReplaceAllString(strings.TrimSpace(userNameRaw), "$1"), nil
}

// Funcs

func (m *NavbarModule) IsLoggedIn() (bool, error) {
	return m.userLink().Visible()
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
