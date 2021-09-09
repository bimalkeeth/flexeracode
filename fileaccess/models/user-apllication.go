package models

// UserApplication data object to hold license related
type UserApplication struct {
	ComputerID    string `json:"computerId"`
	UserID        string `json:"userId"`
	ApplicationID string `json:"applicationId"`
	ComputerType  string `json:"computerType"`
	Comment       string `json:"comment"`
}
