package models

import (
	"fmt"

	"github.com/gky360/atsrv/constants"
)

type Contest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Contest) Host() string {
	return c.ID + ".contests." + constants.AtCoderHost
}
