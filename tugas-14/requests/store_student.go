package requests

type StoreStudentRequest struct {
	Name   string `json:"name"`
	Grade  uint   `json:"grade"`
	Subject string `json:"subject"`
}