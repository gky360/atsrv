package models

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Contest struct {
	ID   string `json:"id" query:"id" yaml:"id"`
	Name string `json:"name" query:"name" yaml:"name"`
}

func (c *Contest) ToYaml() (string, error) {
	d, err := yaml.Marshal(&c)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}
