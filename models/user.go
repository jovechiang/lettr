package models

type User struct {
	UserID    string `json:"userID"`
	PhoneNum  string `json:"phoneNum"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
