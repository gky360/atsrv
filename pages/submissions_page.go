package pages

import (
	"fmt"
	"net/url"
	"path"
	"strconv"

	"github.com/gky360/atsrv/models"
	"github.com/sclevine/agouti"
)

type SubmissionsPage struct {
	page      *agouti.Page
	contestID string
	taskID    string
	lang      models.Language
}

func (p *SubmissionsPage) Page() *agouti.Page {
	return p.page
}

func (p *SubmissionsPage) TargetPath() string {
	contestPath := "/contests/" + p.contestID + "/submissions/me"
	u := &url.URL{Path: contestPath}
	q := u.Query()
	if p.taskID != "" {
		q.Set("f.Task", p.taskID)
	}
	if p.lang != models.LangNone {
		q.Set("f.Language", strconv.Itoa(p.lang.Int()))
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewSubmissionsPage(page *agouti.Page, contestID, taskID string, lang models.Language) (*SubmissionsPage, error) {
	p := &SubmissionsPage{
		page,
		contestID,
		taskID,
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
		if colsCnt, _ := sbmCols.Count(); colsCnt != 10 {
			return nil, fmt.Errorf("found invalid element")
		}
		sbmIDHref, err := sbmCols.At(9).Find("a").Attribute("href")
		if err != nil {
			return nil, err
		}
		sbmID, _ := strconv.Atoi(path.Base(sbmIDHref))

		sbms[i] = &models.Submission{
			ID:           sbmID,
			Lang:         models.NewLanguage(selectionToStr(sbmCols.At(3))),
			Score:        selectionToInt(sbmCols.At(4)),
			SourceLength: selectionToInt(sbmCols.At(5)),
			Status:       selectionToStr(sbmCols.At(6)),
			Time:         selectionToInt(sbmCols.At(7)),
			Memory:       selectionToInt(sbmCols.At(8)),
			CreatedAt:    selectionToStr(sbmCols.At(0)),
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
