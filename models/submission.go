package models

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Submission struct {
	ID           int      `json:"id" yaml:"id"`
	Source       string   `json:"source" yaml:"source,omitempty"`
	Lang         Language `json:"lang" yaml:"lang"`
	Score        int      `json:"score" yaml:"score"`
	SourceLength int      `json:"source_length" yaml:"source_length"`
	Status       string   `json:"status" yaml:"status"`
	Time         int      `json:"time" yaml:"time"`
	Memory       int      `json:"memory" yaml:"memory"`
	CreatedAt    string   `json:"created_at" yaml:"created_at"`
}

func (sbm *Submission) ToYaml() (string, error) {
	d, err := yaml.Marshal(&sbm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func (_sbm *Submission) ToYamlShort() (string, error) {
	sbm := _sbm
	sbm.Source = ""
	return sbm.ToYaml()
}

func SubmissionsToYaml(sbms []*Submission) (string, error) {
	d, err := yaml.Marshal(&sbms)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func SubmissionsToYamlShort(_sbms []*Submission) (string, error) {
	sbms := make([]*Submission, len(_sbms))
	copy(sbms, _sbms)
	for i := range sbms {
		sbms[i].Source = ""
	}
	return SubmissionsToYaml(sbms)
}
