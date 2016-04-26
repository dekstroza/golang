package models

//ApplicationUser describes single entry in that table
type ApplicationUser struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
