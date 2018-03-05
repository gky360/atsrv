package models

type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	FullName string `json:"full_name"`
}
