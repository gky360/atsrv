package pages

import (
	"fmt"
	"regexp"
	"strconv"
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

func (p *TaskPage) TargetPath() string {
	return "/contests/" + p.contestID + "/tasks/" + p.taskID + "?lang=ja"
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
	const JaSelector = selector + " > .lang .lang-ja"
	sel := p.page.Find(JaSelector)
	if _, err := sel.Count(); err != nil {
		// element not found
		// this is for old contests
		sel = p.page.All(selector).At(0)
	}
	return sel
}

func (p *TaskPage) scoreSpan() *agouti.Selection {
	const xpath = ".//p[contains(text(), '配点')]/var"
	return p.statement().FindByXPath(xpath)
}

func (p *TaskPage) samplePres() *agouti.MultiSelection {
	const xpath = ".//*[contains(@class, 'part')]//section//pre[contains(@id, 'pre-sample')]"
	return p.statement().AllByXPath(xpath)
}

func (p *TaskPage) samplePre(index int) *agouti.Selection {
	selector := fmt.Sprintf(".part section pre#pre-sample%d", index)
	return p.statement().Find(selector)
}

// Values

func (p *TaskPage) taskNameAndTitle() (string, string, error) {
	nameAndTitleRaw, err := p.titleH2().Text()
	if err != nil {
		return "", "", err
	}
	nameAndTitle := strings.SplitN(nameAndTitleRaw, " - ", 2)
	return strings.TrimSpace(nameAndTitle[0]), strings.TrimSpace(nameAndTitle[1]), nil
}

func (p *TaskPage) taskScore() (int, error) {
	scoreRaw, err := p.scoreSpan().Text()
	if err != nil {
		return 0, nil
	}
	rep := regexp.MustCompile("[0-9]+")
	if err != nil {
		return 0, err
	}
	score, _ := strconv.Atoi(rep.FindString(scoreRaw))
	return score, nil
}

func (p *TaskPage) sample(num int) (*models.Sample, error) {
	inputStr, err := p.samplePre(num*2 - 2).Text()
	if err != nil {
		return nil, err
	}
	outputStr, err := p.samplePre(num*2 - 1).Text()
	if err != nil {
		return nil, err
	}

	return &models.Sample{
		Num:    num,
		Input:  inputStr,
		Output: outputStr,
	}, nil
}

func (p *TaskPage) samples() ([]*models.Sample, error) {
	presCnt, err := p.samplePres().Count()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("found %d samples\n", presCnt/2)
	samples := make([]*models.Sample, presCnt/2)
	for i := range samples {
		samples[i], err = p.sample(i + 1)
		if err != nil {
			return nil, err
		}
	}

	return samples, nil
}

// Funcs

func (p *TaskPage) GetTask() (*models.Task, error) {
	taskName, taskTitle, err := p.taskNameAndTitle()
	if err != nil {
		return nil, err
	}
	taskScore, err := p.taskScore()
	if err != nil {
		return nil, err
	}
	samples, err := p.samples()
	if err != nil {
		fmt.Println(err)
	}
	return &models.Task{
		ID:      p.taskID,
		Name:    taskName,
		Title:   taskTitle,
		Score:   taskScore,
		Samples: samples,
	}, nil
}
