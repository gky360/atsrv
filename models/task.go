package models

type Task struct {
	ID    string `json:"id" yaml:"id"`
	Name  string `json:"name" yaml:"name"`
	Title string `json:"title" yaml:"title"`
	Score int    `json:"score" yaml:"score"`
}
