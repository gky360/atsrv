package models

type Submission struct {
	ID     int    `json:"id" yaml:"id"`
	Source string `json:"source" yaml:"source"`
}
