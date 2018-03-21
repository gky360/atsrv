package models

import (
	"fmt"

	"github.com/gky360/atsrv/constants"
	"gopkg.in/yaml.v2"
)

type Contest struct {
	ID   string `json:"id" query:"id" yaml:"id"`
	Name string `json:"name" query:"name" yaml:"name"`
}

func (c *Contest) Host() string {
	return c.ID + ".contest." + constants.AtCoderHost
}

func (c *Contest) URL() string {
	return "https://" + c.Host()
}

func (c *Contest) ToYaml() (string, error) {
	d, err := yaml.Marshal(&c)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s", string(d)), nil
}

func GetCurrentContest() (Contest, error) {
	contest := Contest{
		ID:   "agc021",
		Name: "AtCoder Grand Contest 021",
	}
	return contest, nil
}
