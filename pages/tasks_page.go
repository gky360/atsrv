package pages

import (
	"github.com/sclevine/agouti"
)

const targetPath string = "/assignments"

type TasksPage struct {
	page *agouti.Page
}

func NewTasksPage(page *agouti.Page) (Page, error) {
	p := new(TasksPage)
	p.page = page
	if err := At(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *TasksPage) GetTargetPath() string {
	return targetPath
}

func (p *TasksPage) GetPage() *agouti.Page {
	return p.page
}
