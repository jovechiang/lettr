package models

type LoginResponse struct {
	Token      string `json:"token"`
	LoggedUser *User  `json:"loggedUser"`
}
