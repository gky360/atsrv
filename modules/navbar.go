package modules

import (
	"github.com/sclevine/agouti"
)

type NavbarModule struct {
	page *agouti.Page
}

func NewNavbarModule(page *agouti.Page) *NavbarModule {
	return &NavbarModule{page: page}
}
