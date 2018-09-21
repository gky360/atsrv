package pages

import (
	"fmt"
	"net/url"
	"path"
	"strconv"

	"github.com/sclevine/agouti"

	"github.com/gky360/atsrv/constants"
	"github.com/gky360/atsrv/models"
)

type SubmissionsPage struct {
	page      *agouti.Page
	contestID string
	taskID    string
	status    string
	lang      models.Language
}

func (p *SubmissionsPage) Page() *agouti.Page {
	return p.page
}

func (p *SubmissionsPage) Hostname() string {
	return constants.AtCoderHost
}

func (p *SubmissionsPage) TargetPath() string {
	sbmsPath := "/contests/" + p.contestID + "/submissions/me"
	u := &url.URL{Path: sbmsPath}
	q := u.Query()
	if p.taskID != "" {
		q.Set("f.Task", p.taskID)
	}
	if p.status != "" {
		q.Set("f.Status", p.status)
	}
	if p.lang != models.LangNone {
		q.Set("f.Language", strconv.Itoa(p.lang.Int()))
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewSubmissionsPage(page *agouti.Page, contestID, taskID, status string, lang models.Language) (*SubmissionsPage, error) {
	p := &SubmissionsPage{
		page,
		contestID,
		taskID,
		status,
		lang,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *SubmissionsPage) sbmsTable() *agouti.Selection {
	const selector = "#main-container .panel-submission .table"
	return p.page.Find(selector)
}

func (p *SubmissionsPage) sbmTRs() *agouti.MultiSelection {
	const selector = "tbody tr"
	return p.sbmsTable().All(selector)
}

func (p *SubmissionsPage) sbmCols(index int) *agouti.MultiSelection {
	return p.sbmTRs().At(index).All("td")
}

// Values

func (p *SubmissionsPage) sbms() ([]*models.Submission, error) {
	cnt, _ := p.sbmTRs().Count()
	sbms := make([]*models.Submission, cnt)
	for i := range sbms {
		sbmCols := p.sbmCols(i)
		isWJ := false
		colsCnt, _ := sbmCols.Count()
		switch colsCnt {
		case 8:
			isWJ = true
		case 10:
			isWJ = false
		default:
			return nil, fmt.Errorf("found invalid element")
		}
		sbmIDHref, err := sbmCols.At(colsCnt - 1).Find("a").Attribute("href")
		if err != nil {
			return nil, err
		}
		sbmID, _ := strconv.Atoi(path.Base(sbmIDHref))
		taskIDHref, err := sbmCols.At(1).Find("a").Attribute("href")
		if err != nil {
			return nil, err
		}
		taskID := path.Base(taskIDHref)

		sbms[i] = &models.Submission{
			ID:           sbmID,
			Lang:         models.NewLanguage(selectionToStr(sbmCols.At(3))),
			Score:        selectionToInt(sbmCols.At(4)),
			SourceLength: selectionToInt(sbmCols.At(5)),
			Status:       selectionToStr(sbmCols.At(6).Find("span")),
			CreatedAt:    selectionToStr(sbmCols.At(0)),
			Task: models.NewTaskWithFullName(
				taskID,
				selectionToStr(sbmCols.At(1)),
			),
		}
		if !isWJ {
			sbms[i].Time = selectionToInt(sbmCols.At(7))
			sbms[i].Memory = selectionToInt(sbmCols.At(8))
		}
	}
	return sbms, nil
}

// Funcs

func (p *SubmissionsPage) GetSubmissions() ([]*models.Submission, error) {
	return p.sbms()
}

var ErrSubmissionNotFound = fmt.Errorf("submission not found")

func (p *SubmissionsPage) GetSubmission(sbmID int) (*models.Submission, error) {
	sbms, err := p.sbms()
	if err != nil {
		return nil, err
	}
	for _, sbm := range sbms {
		if sbm.ID == sbmID {
			return sbm, nil
		}
	}
	return nil, ErrSubmissionNotFound
}
