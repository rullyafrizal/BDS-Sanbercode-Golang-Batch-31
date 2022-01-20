package models

type Student struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Grade   uint   `json:"grade"`
	Index   string `json:"index"`
}
