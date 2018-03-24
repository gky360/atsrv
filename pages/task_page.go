package pages

import (
	// "fmt"
	// "strings"

	// "github.com/gky360/atsrv/models"
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

func (p *TaskPage) TargetPath() string {
	return "/contests/" + p.contestID + "/tasks/" + p.taskID
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
	const selector = "#main-container .h2"
	return p.page.Find(selector)
}

func (p *TaskPage) statement() *agouti.Selection {
	const selector = "#main-container #task-statement"
	const JaSelector = selector + ".lang .lang-ja"
	sel := p.page.Find(JaSelector)
	if _, err := sel.Count(); err != nil {
		// element not found
		// this is for old contests
		sel = p.page.Find(selector)
	}
	return sel
}

// Values

// Funcs
