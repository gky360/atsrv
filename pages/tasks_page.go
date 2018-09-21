package pages

import (
	"fmt"
	"path"
	"strings"

	"github.com/sclevine/agouti"

	"github.com/gky360/atsrv/constants"
	"github.com/gky360/atsrv/models"
)

type TasksPage struct {
	page      *agouti.Page
	contestID string
}

func (p *TasksPage) Page() *agouti.Page {
	return p.page
}

func (p *TasksPage) Hostname() string {
	return constants.AtCoderHost
}

func (p *TasksPage) TargetPath() string {
	return "/contests/" + p.contestID + "/tasks"
}

func NewTasksPage(page *agouti.Page, contestID string) (*TasksPage, error) {
	p := &TasksPage{
		page:      page,
		contestID: contestID,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *TasksPage) tasksTable() *agouti.Selection {
	const selector = "#main-container .table"
	return p.page.Find(selector)
}

func (p *TasksPage) taskTRs() *agouti.MultiSelection {
	const selector = "tbody tr"
	return p.tasksTable().All(selector)
}

func (p *TasksPage) taskCols(index int) *agouti.MultiSelection {
	return p.taskTRs().At(index).All("td")
}

// Values

func (p *TasksPage) tasks() ([]*models.Task, error) {
	cnt, _ := p.taskTRs().Count()
	tasks := make([]*models.Task, cnt)
	for i := range tasks {
		taskCols := p.taskCols(i)
		if colsCnt, _ := taskCols.Count(); colsCnt != 5 {
			return nil, fmt.Errorf("found invalid element")
		}
		taskIDHref, err := taskCols.At(0).Find("a").Attribute("href")
		if err != nil {
			return nil, err
		}
		taskNameRaw, _ := taskCols.At(0).Text()
		taskTitleRaw, _ := taskCols.At(1).Text()
		tasks[i] = &models.Task{
			ID:    path.Base(taskIDHref),
			Name:  strings.TrimSpace(taskNameRaw),
			Title: strings.TrimSpace(taskTitleRaw),
		}
	}
	return tasks, nil
}

// Funcs

func (p *TasksPage) GetTasks() ([]*models.Task, error) {
	return p.tasks()
}

var ErrTaskNameNotFound = fmt.Errorf("task name not found")

func (p *TasksPage) GetTaskID(taskName string) (string, error) {
	cnt, _ := p.taskTRs().Count()
	for i := 0; i < cnt; i++ {
		taskCols := p.taskCols(i)
		taskNameRaw, _ := taskCols.At(0).Text()
		if strings.EqualFold(taskName, strings.TrimSpace(taskNameRaw)) {
			taskIDHref, err := taskCols.At(0).Find("a").Attribute("href")
			if err != nil {
				return "", err
			}
			return path.Base(taskIDHref), nil
		}
	}
	return "", ErrTaskNameNotFound
}
