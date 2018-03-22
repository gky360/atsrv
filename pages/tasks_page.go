package pages

import (
	"github.com/gky360/atsrv/modules"
	"github.com/sclevine/agouti"
)

type TasksPage struct {
	page      *agouti.Page
	contestID string
	navbar    *modules.NavbarModule
}

func (p *TasksPage) Page() *agouti.Page {
	return p.page
}

func (p *TasksPage) TargetHost() string {
	return ContestHost(p.contestID)
}

func (p *TasksPage) TargetPath() string {
	return "/assignments"
}

func NewTasksPage(page *agouti.Page, contestID string) (*TasksPage, error) {
	p := &TasksPage{
		page:      page,
		contestID: contestID,
		navbar:    modules.NewNavbarModule(page),
	}
	if err := At(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *TasksPage) Login(userID, password string) error {
	return p.navbar.Login(userID, password)
}
