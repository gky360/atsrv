package pages

import (
	"fmt"
	"strings"

	"github.com/gky360/atsrv/models"
	"github.com/sclevine/agouti"
)

type TasksPage struct {
	page      *agouti.Page
	contestID string
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
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *TasksPage) tasksTable() *agouti.Selection {
	const selector = "#outer-inner .table"
	return p.page.Find(selector)
}

func (p *TasksPage) taskTRs() *agouti.MultiSelection {
	const selector = "tbody tr"
	return p.tasksTable().All(selector)
}

// Values

func (p *TasksPage) tasks() ([]models.Task, error) {
	tasksTRs := p.taskTRs()
	cnt, _ := tasksTRs.Count()
	tasks := make([]models.Task, cnt)
	for i := 0; i < cnt; i++ {
		taskCols := tasksTRs.At(i).All("td")
		if colsCnt, _ := taskCols.Count(); colsCnt != 5 {
			return nil, fmt.Errorf("found invalid element")
		}
		taskNameRaw, _ := taskCols.At(0).Text()
		taskTitleRaw, _ := taskCols.At(1).Text()
		tasks[i].ID = "hoge"
		tasks[i].Name = strings.TrimSpace(taskNameRaw)
		tasks[i].Title = strings.TrimSpace(taskTitleRaw)
	}
	return tasks, nil
}

// Funcs

func (p *TasksPage) GetTasks() ([]models.Task, error) {
	return p.tasks()
}
