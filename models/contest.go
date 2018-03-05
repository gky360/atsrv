package models

import (
	"github.com/gky360/atsrv/constants"
)

type Contest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Contest) Host() string {
	return c.ID + ".contest." + constants.AtCoderHost
}

func (c *Contest) URL() string {
	return "https://" + c.Host()
}

func GetCurrentContest() (Contest, error) {
	contest := Contest{
		ID:   "agc021",
		Name: "AtCoder Grand Contest 021",
	}
	return contest, nil
}
