package models

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Submission struct {
	ID     int    `json:"id" yaml:"id"`
	Source string `json:"source" yaml:"source,omitempty"`
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
	d, err := yaml.Marshal(&sbm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func SubmissionsToYaml(sbms []Submission) (string, error) {
	d, err := yaml.Marshal(&sbms)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func SubmissionsToYamlShort(_sbms []Submission) (string, error) {
	sbms := _sbms
	for i := range sbms {
		sbms[i].Source = ""
	}

	d, err := yaml.Marshal(&sbms)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}
