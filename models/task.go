package models

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type Sample struct {
	Num    int    `json:"num" yaml:"num"`
	Input  string `json:"input" yaml:"input"`
	Output string `json:"output" yaml:"output"`
}

type Task struct {
	ID      string    `json:"id" yaml:"id"`
	Name    string    `json:"name" yaml:"name"`
	Title   string    `json:"title" yaml:"title"`
	Score   int       `json:"score,omitempty" yaml:"score,omitempty"`
	Samples []*Sample `json:"samples,omitempty" yaml:"samples,omitempty"`
}

func NewTaskWithFullName(taskID, fullName string) *Task {
	nameAndTitle := strings.SplitN(fullName, " - ", 2)
	if len(nameAndTitle) != 2 {
		return &Task{Title: fullName}
	}
	return &Task{
		ID:    taskID,
		Name:  nameAndTitle[0],
		Title: nameAndTitle[1],
	}
}

func (task *Task) ToYaml() (string, error) {
	d, err := yaml.Marshal(&task)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func (_task *Task) ToYamlShort() (string, error) {
	task := _task
	task.Samples = nil
	return task.ToYaml()
}

func TasksToYaml(tasks []*Task) (string, error) {
	d, err := yaml.Marshal(&tasks)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func TasksToYamlShort(_tasks []*Task) (string, error) {
	tasks := make([]*Task, len(_tasks))
	copy(tasks, _tasks)
	for i := range tasks {
		tasks[i].Samples = nil
	}
	return TasksToYaml(tasks)
}
