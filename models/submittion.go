package models

type Submittion struct {
	ID     int    `json:"id"`
	TaskID string `json:"task_id"`
	Source string `json:"source"`
}
