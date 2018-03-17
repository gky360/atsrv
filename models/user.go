package models

type User struct {
	ID       string `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty" yaml:"token"`
}
