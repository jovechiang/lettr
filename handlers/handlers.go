package handlers

import (
	. "github.com/jovechiang/lettr/models"
)

func ProcessLogin(phoneNum string, password string) *LoginResponse {
	return &LoginResponse{
		Token: "dummyToken",
		LoggedUser: &User{
			UserID:    "1",
			PhoneNum:  "+19178687749",
			FirstName: "Kan",
			LastName:  "Uber",
		},
	}
}
