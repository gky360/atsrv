package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}
