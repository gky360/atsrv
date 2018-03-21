package models

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type User struct {
	ID       string `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty" yaml:"token"`
}

func (user *User) ToYaml() (string, error) {
	d, err := yaml.Marshal(&user)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}
