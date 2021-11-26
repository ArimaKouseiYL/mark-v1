package models

type DbConfig struct {
	Url      string `json:"url"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
