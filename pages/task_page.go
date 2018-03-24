package pages

import (
	"fmt"
	"strings"

	"github.com/gky360/atsrv/models"
	"github.com/sclevine/agouti"
)

type TaskPage struct {
	page      *agouti.Page
	contestID string
	taskID    string
}

func (p *TaskPage) Page() *agouti.Page {
	return p.page
}

func (p *TaskPage) TargetHost() string {
	return ContestHost(p.contestID)
}

func (p *TaskPage) TargetPath() string {
	return "/tasks/" + p.taskID
}

func NewTaskPage(page *agouti.Page, contestID string, taskID string) (*TaskPage, error) {
	p := &TaskPage{
		page:      page,
		contestID: contestID,
		taskID:    taskID,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *TaskPage) titleH2() *agouti.Selection {
	const selector = "#outer-inner h2"
	return p.page.Find(selector)
}

func (p *TaskPage) statement() *agouti.Selection {
	const selector = "#outer-inner > #task-statement > .lang > .lang-en"
	return p.page.Find(selector)
}

// Values

// Funcs
