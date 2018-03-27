package pages

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/gky360/atsrv/models"
	"github.com/sclevine/agouti"
)

type SubmissionsPage struct {
	page      *agouti.Page
	contestID string
	taskID    string
	Lang      models.Language
}

func (p *SubmissionsPage) Page() *agouti.Page {
	return p.page
}

func (p *SubmissionsPage) TargetPath() string {
	return "/contests/" + p.contestID + "/submissions/me"
}

func NewSubmissionsPage(page *agouti.Page, contestID string) (*SubmissionsPage, error) {
	p := &SubmissionsPage{
		page:      page,
		contestID: contestID,
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

func selectionToStr(sel *agouti.Selection) string {
	raw, _ := sel.Text()
	return strings.TrimSpace(raw)
}

func selectionToInt(sel *agouti.Selection) int {
	ret, _ := strconv.Atoi(selectionToStr(sel))
	return ret
}

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
