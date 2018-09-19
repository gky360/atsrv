package pages

import (
	"fmt"
	"strings"

	"github.com/gky360/atsrv/constants"
	"github.com/sclevine/agouti"
)

type TestcasesPage struct {
	page *agouti.Page
}

func (p *TestcasesPage) Page() *agouti.Page {
	return p.page
}

func (p *TestcasesPage) Hostname() string {
	return constants.DropboxHost
}

func (p *TestcasesPage) TargetPath() string {
	return constants.AtcoderTestcasesRootPath
}

func NewTestcasesPage(page *agouti.Page) (*TestcasesPage, error) {
	p := &TestcasesPage{
		page: page,
	}
	if err := To(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Elements

func (p *TestcasesPage) foldersBody() *agouti.Selection {
	const selector = ".sl-grid-body"
	return p.page.Find(selector)
}

func (p *TestcasesPage) folderLinks() *agouti.MultiSelection {
	const selector = "li.sl-grid-cell a.sl-link--folder"
	return p.foldersBody().All(selector)
}

// Values

var ErrTestcasesFolderNotFound = fmt.Errorf("contest testcases folder not found")

func (p *TestcasesPage) contestFolderURL(contestID string) (string, error) {
	normContestID := normalizeContestTestcasesFolderName(contestID)
	cnt, _ := p.folderLinks().Count()
	for i := 0; i < cnt; i++ {
		folderLink := p.folderLinks().At(i)
		folderURL, _ := folderLink.Attribute("href")
		folderName, _ := folderLink.Text()
		normFolderName := normalizeContestTestcasesFolderName(folderName)
		if strings.Contains(normFolderName, normContestID) {
			return folderURL, nil
		}
	}

	return "", ErrTestcasesFolderNotFound
}

// Funcs

func normalizeContestTestcasesFolderName(name string) string {
	r := strings.NewReplacer("-", "", "_", "")
	return strings.ToLower(r.Replace(name))
}

func (p *TestcasesPage) GetContestFolderURL(contestID string) (string, error) {
	return p.contestFolderURL(contestID)
}
