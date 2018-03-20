package models

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Task struct {
	ID    string `json:"id" yaml:"id"`
	Name  string `json:"name" yaml:"name"`
	Title string `json:"title" yaml:"title"`
	Score int    `json:"score" yaml:"score"`
}

func (task *Task) ToYaml() (string, error) {
	d, err := yaml.Marshal(&task)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func TasksToYaml(tasks []Task) (string, error) {
	d, err := yaml.Marshal(&tasks)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}
