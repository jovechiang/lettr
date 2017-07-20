package models

type Chat struct {
	ChatID   string
	IsRead   bool
	Chatter  User
	Messages []Message
}
